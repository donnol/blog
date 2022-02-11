---
author: "jdlau"
date: 2022-02-11
linktitle: go runtime chan
menu:
next:
prev:
title: go runtime chan
weight: 10
categories: ['go', 'runtime', 'chan']
tags: ['chan']
---

`src/runtime/chan.go`:

```go
// Invariants:
//  At least one of c.sendq and c.recvq is empty,
//  except for the case of an unbuffered channel with a single goroutine
//  blocked on it for both sending and receiving using a select statement,
//  in which case the length of c.sendq and c.recvq is limited only by the
//  size of the select statement.
//
// For buffered channels, also:
//  c.qcount > 0 implies that c.recvq is empty.
//  c.qcount < c.dataqsiz implies that c.sendq is empty.

// 在文件开头，说明了几个不变量：
//  c.sendq和c.recvq中至少有一个是空的，
//  除非，一个无缓冲管道在一个goroutine里阻塞了，这个管道的发送和接收都使用了一个select语句，这时
//  c.sendq和c.recvq的长度被select语句限制。
// 
// 对于缓冲管道，同样地：
//  c.qcount > 0 表明c.recvq是空的。
//  c.qcount < c.dataqsiz 表明c.sendq是空的。

// 实际的chan类型
type hchan struct {
	qcount   uint           // total data in the queue - 队列里的数据总数量
	dataqsiz uint           // size of the circular queue - 循环队列的大小，make时传进来的值
	buf      unsafe.Pointer // points to an array of dataqsiz elements - dataqsiz元素组成的数组的指针
	elemsize uint16 // 元素大小
	closed   uint32 // 是否关闭
	elemtype *_type // element type - 元素类型
	sendx    uint   // send index - 发送索引
	recvx    uint   // receive index - 接收索引
	recvq    waitq  // list of recv waiters - 等待接收者列表，表明这个管道的接收者；一个链表，里面的每个元素代表一个g；
	sendq    waitq  // list of send waiters - 等待发送者列表，编码这个管道的发送者

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex // 保护chan里的所有字段，以及阻塞在本管道里的sudog；当持有这个锁时，不要改变其它G的状态，因为在栈收缩时可能引起死锁。
}

type waitq struct {
	first *sudog
	last  *sudog
}

// sudog represents a g in a wait list, such as for sending/receiving
// on a channel. - 代表了一个在等待列表的g
//
// sudog is necessary because the g ↔ synchronization object relation
// is many-to-many. A g can be on many wait lists, so there may be
// many sudogs for one g; and many gs may be waiting on the same
// synchronization object, so there may be many sudogs for one object.
// - sudog是必须的，因为g和同步对象关系是多对多。一个g可以在多个等待列表里，因此一个g对应有多个sudog；
// 多个g可以等待同一个同步对象，因此一个对象会对应多个sudog。
//
// sudogs are allocated from a special pool. Use acquireSudog and
// releaseSudog to allocate and free them.
// - sudog从一个特殊池子里分配，使用acquireSudog分配和releaseSudog释放它们。
type sudog struct {
	// The following fields are protected by the hchan.lock of the
	// channel this sudog is blocking on. shrinkstack depends on
	// this for sudogs involved in channel ops.
    // - 以下字段由hchan.lock来保护。

	g *g // 代表的g

	next *sudog // 链表中的下一个
	prev *sudog // 链表中的上一个
	elem unsafe.Pointer // data element (may point to stack) - 数据元素，可能是指向栈的指针

	// The following fields are never accessed concurrently.
	// For channels, waitlink is only accessed by g.
	// For semaphores, all fields (including the ones above)
	// are only accessed when holding a semaRoot lock.
    // - 以下字段永远不会被并发访问。
    // 对于管道，waitlink只会被g访问。
    // 对于信号量，所有字段（包括上面的）只有在持有semaRoot锁时才能被访问

	acquiretime int64 // 获取时间
	releasetime int64 // 释放时间
	ticket      uint32 // 票据

	// isSelect indicates g is participating in a select, so
	// g.selectDone must be CAS'd to win the wake-up race.
	isSelect bool // 表明g是否参与到了一个select里，从而使得g.selectDone必须CAS地去赢得唤醒竞赛

	// success indicates whether communication over channel c
	// succeeded. It is true if the goroutine was awoken because a
	// value was delivered over channel c, and false if awoken
	// because c was closed.
	success bool // 表明管道的通信是否成功了，如果goroutine因为一个值被管道传送到来而唤醒即为成功

	parent   *sudog // semaRoot binary tree - 根信号量二叉树
	waitlink *sudog // g.waiting list or semaRoot - g的等待列表或semaRoot
	waittail *sudog // semaRoot
	c        *hchan // channel - 所属管道
}

// 新建
func makechan(t *chantype, size int) *hchan {
	elem := t.elem

	// compiler checks this but be safe.
	if elem.size >= 1<<16 { // 管道的元素大小不能太大
		throw("makechan: invalid channel element type")
	}
    // const hchanSize uintptr = 96
	if hchanSize%maxAlign != 0 || elem.align > maxAlign { // 对齐检查
		throw("makechan: bad alignment")
	}

    // 元素大小乘以管道大小，计算出来所需内存大小
	mem, overflow := math.MulUintptr(elem.size, uintptr(size)) 
	if overflow || mem > maxAlloc-hchanSize || size < 0 {
		panic(plainError("makechan: size out of range"))
	}

	// Hchan does not contain pointers interesting for GC when elements stored in buf do not contain pointers.
	// buf points into the same allocation, elemtype is persistent.
	// SudoG's are referenced from their owning thread so they can't be collected.
	// TODO(dvyukov,rlh): Rethink when collector can move allocated objects.
	var c *hchan
	switch {
	case mem == 0:
		// Queue or element size is zero.
		c = (*hchan)(mallocgc(hchanSize, nil, true))
		// Race detector uses this location for synchronization.
		c.buf = c.raceaddr()
	case elem.ptrdata == 0:
		// Elements do not contain pointers. -- 元素没有包含指针
		// Allocate hchan and buf in one call.
		c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
		c.buf = add(unsafe.Pointer(c), hchanSize)
	default:
		// Elements contain pointers. -- 元素包含指针
		c = new(hchan)
		c.buf = mallocgc(mem, elem, true)
	}

	c.elemsize = uint16(elem.size)
	c.elemtype = elem
	c.dataqsiz = uint(size)
	lockInit(&c.lock, lockRankHchan) // 初始化锁

	if debugChan {
		print("makechan: chan=", c, "; elemsize=", elem.size, "; dataqsiz=", size, "\n")
	}
	return c
}

// 发送
func sendDirect(t *_type, sg *sudog, src unsafe.Pointer) {
	// src is on our stack, dst is a slot on another stack.
    // - src是在我们的栈上，dst是另一个栈上的槽

	// Once we read sg.elem out of sg, it will no longer
	// be updated if the destination's stack gets copied (shrunk).
	// So make sure that no preemption points can happen between read & use.
	dst := sg.elem
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	// No need for cgo write barrier checks because dst is always
	// Go memory.
	memmove(dst, src, t.size) // 移动src到dst
}

// 接收 -- 请看源码

// 关闭
func closechan(c *hchan) {
	if c == nil {
		panic(plainError("close of nil channel"))
	}

	lock(&c.lock)
	if c.closed != 0 { // 已关闭的chan，如果再次关闭会panic
		unlock(&c.lock)
		panic(plainError("close of closed channel"))
	}

	if raceenabled {
		callerpc := getcallerpc()
		racewritepc(c.raceaddr(), callerpc, abi.FuncPCABIInternal(closechan))
		racerelease(c.raceaddr())
	}

	c.closed = 1 // 设为关闭

	var glist gList

    // 先释放接收者，再释放发送者

	// release all readers
	for {
		sg := c.recvq.dequeue() // 逐个出队sudog
		if sg == nil {
			break
		}
		if sg.elem != nil {
			typedmemclr(c.elemtype, sg.elem) // 清理元素
			sg.elem = nil
		}
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = unsafe.Pointer(sg)
		sg.success = false
		if raceenabled {
			raceacquireg(gp, c.raceaddr())
		}
		glist.push(gp) // 把关联的g存到glist里
	}

	// release all writers (they will panic)
	for {
		sg := c.sendq.dequeue()
		if sg == nil {
			break
		}
		sg.elem = nil
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = unsafe.Pointer(sg)
		sg.success = false
		if raceenabled {
			raceacquireg(gp, c.raceaddr())
		}
		glist.push(gp)
	}
	unlock(&c.lock)

	// Ready all Gs now that we've dropped the channel lock.
	for !glist.empty() {
		gp := glist.pop() // 逐个处理g
		gp.schedlink = 0
		goready(gp, 3) // 因为我们已经释放了这些g所关联的chan，所以让这些g进入ready状态，准备运行 -- Mark gp ready to run.
	}
}
```

`src/runtime/type.go`:

```go
// Needs to be in sync with ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/reflectdata/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
// ../internal/reflectlite/type.go:/^type.rtype.
// 
// 这个类型必须与链接器、编译器、反射等地方的类型保持同步
type _type struct {
	size       uintptr // 大小
	ptrdata    uintptr // size of memory prefix holding all pointers - 持有所有指针的内存前缀大小
	hash       uint32 // 哈希值
	tflag      tflag // 类型标记
	align      uint8 // 对齐
	fieldAlign uint8 // 字段对齐
	kind       uint8 // 种类
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool // 比较本类型的两个对象的指针的方法
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte // 存储了垃圾收集器所需的GC类型数据
	str       nameOff // 名称偏移
	ptrToThis typeOff // 类型偏移
}
```
