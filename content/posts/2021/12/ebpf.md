---
author: "jdlau"
date: 2021-12-24
linktitle: ebpf
menu:
next:
prev:
title: ebpf
weight: 10
categories: ['linux']
tags: ['ebpf']
---

[ebpf](https://ebpf.io/): 扩展伯克利包过滤器。

下面的内容主要来源于[译文](https://mp.weixin.qq.com/s/xxBQuj-iD103kWeicanUFA)。

## 用处

> 目前，主要有两大组触发器。
>
> 第一组用于处理网络数据包和管理网络流量。它们是 XDP、流量控制事件及其他几个事件。
>
> 以下情况需要用到这些事件：
>
>> 创建简单但非常有效的防火墙。Cloudflare 和 Facebook 等公司使用 BPF 程序来过滤掉大量的寄生流量，并打击最大规模的 DDoS 攻击。由于处理发生在数据包生命的最早阶段，直接在内核中进行（BPF 程序的处理有时甚至可以直接推送到网卡中进行），因此可以通过这种方式处理巨量的流量。这些事情过去都是在专门的网络硬件上完成的。
>>
>> 创建更智能、更有针对性、但性能更好的防火墙——这些防火墙可以检查通过的流量是否符合公司的规则、是否存在漏洞模式等。例如，Facebook 在内部进行这种审计，而一些项目则对外销售这类产品。
>>
>> 创建智能负载均衡器。最突出的例子就是 Cilium 项目，它最常被用作 K8s 集群中的网格网络。Cilium 对流量进行管理、均衡、重定向和分析。所有这些都是在内核运行的小型 BPF 程序的帮助下完成的，以响应这个或那个与网络数据包或套接字相关的事件。
>
> 这是第一组与网络问题相关并能够影响网络通信行为的触发器。第二组则与更普遍的可观察性相关；在大多数情况下，这组的程序无法影响任何事件，而只能“观察”。这才是我更感兴趣的。
>
> 这组的触发器有如下几个：
>
>> perf 事件（perf events）——与性能和 perf Linux 分析器相关的事件：硬件处理器计数器、中断处理、小 / 大内存异常拦截等等。例如，我们可以设置一个处理程序，每当内核需要从 swap 读取内存页时，该处理程序就会运行。例如，想象有这样一个实用程序，它显示了当前所有使用 swap 的程序。
>>
>> 跟踪点（tracepoints）——内核源代码中的静态（由开发人员定义）位置，通过附加到这些位置，你可以从中提取静态信息（开发人员先前准备的信息）。在这种情况下，静态似乎是一件坏事，因为我说过，日志的缺点之一就是它们只包含了程序员最初放在那里的内容。从某种意义上说，这是正确的，但跟踪点有三个重要的优势：
>>
> 有相当多的跟踪点散落在内核中最有趣的地方
>
> 当它们不“开启”时，它们不使用任何资源
>
> 它们是 API 的一部分，它们是稳定的，不会改变。这非常重要，因为我们将提到的其他触发器缺少稳定的 API。
>
> 例如，假设有一个关于显示的实用程序，内核出于某种原因没有给它时间执行。你坐着纳闷为什么它这么慢，而 pprof 却没有显示任何什么有趣的东西。
>>
>> USDT——与跟踪点相同，但是它适用于用户空间的程序。也就是说，作为程序员，你可以将这些位置添加到你的程序中。并且许多大型且知名的程序和编程语言都已经采用了这些跟踪方法：例如 MySQL、或者 PHP 和 Python 语言。通常，它们的默认设置为“关闭”，如果要打开它们，需要使用 enable-dtrace 参数或类似的参数来重新构建解释器。是的，我们还可以在 Go 中注册这种类跟踪。你可能已经识别出参数名称中的单词 DTrace。关键在于，这些类型的静态跟踪是由 Solaris) 操作系统中诞生的同名系统所推广的。例如，想象一下，何时创建新线程、何时启动 GC 或与特定语言或系统相关的其他内容，我们都能够知道是怎样的一种场景。
>
> 这是另一种魔法开始的地方：
>
>> Ftrace 触发器为我们提供了在内核的任何函数开始时运行 BPF 程序的选项。这是完全动态的。这意味着内核将在你选择的任何内核函数或者在所有内核函数开始执行之前，开始执行之前调用你的 BPF 函数。你可以连接到所有内核函数，并在输出时获取所有调用的有吸引力的可视化效果。
>>
>> kprobes/uprobes 提供的功能与 ftrace 几乎相同，但在内核和用户空间中执行函数时，你可以选择将其附加到任何位置上。如果在函数的中间，变量上有一个“if”，并且能为这个变量建立一个值的直方图，那就不是问题。
>>
>> kretprobes/uretprobes——这里的一切都类似于前面的触发器，但是它们可以在内核函数或用户空间中的函数返回时触发。这类触发器便于查看函数的返回内容以及测量执行所需的时间。例如，你可以找出“fork”系统调用返回的 PID。
>
> 我再重复一遍，所有这些最奇妙之处在于，当我们的 BPF 程序为了响应这些触发器而被调用之后，我们可以很好地“环顾四周”：**读取函数的参数，记录时间，读取变量，读取全局变量，进行堆栈跟踪，保存一些内容以备后用，将数据发送到用户空间进行处理，和 / 或从用户空间获取数据或一些其他控制命令以进行过滤**。简直不可思议！

## 使用

先用`man bpf`看下它是怎么被定义的：

```c
NAME
    bpf - perform a command on an extended BPF map or program
	-- 在bpf map或程序里执行一个命令

SYNOPSIS
    #include <linux/bpf.h>

    int bpf(int cmd, union bpf_attr *attr, unsigned int size);
	-- bpf_attr里有些什么呢？

DESCRIPTION
    The  bpf() system call performs a range of operations related to extended Berkeley Packet
    Filters.  Extended BPF (or eBPF) is similar to the original ("classic") BPF  (cBPF)  used
    to  filter  network packets.  For both cBPF and eBPF programs, the kernel statically ana‐
    lyzes the programs before loading them, in order to ensure that they cannot harm the run‐
    ning system.
	-- 在程序被加载之前分析它，确保它们不会伤害到运行中的系统。

    eBPF  extends cBPF in multiple ways, including the ability to call a fixed set of in-ker‐
    nel helper functions (via the BPF_CALL opcode extension  provided  by  eBPF)  and  access
    shared data structures such as eBPF maps.
	-- ebpf在很多方面扩展了cbpf，包括可以调用固定的钩子函数的能力，可以访问共享的数据结构（如：ebpf maps）。

Extended BPF Design/Architecture
	eBPF  maps  are a generic data structure for storage of different data types.  Data types
	are generally treated as binary blobs, so a user just specifies the size of the  key  and
	the  size of the value at map-creation time.  In other words, a key/value for a given map
	can have an arbitrary structure.

	A user process can create multiple maps (with key/value-pairs being opaque bytes of data)
	and  access  them via file descriptors.  Different eBPF programs can access the same maps
	in parallel.  It's up to the user process and eBPF program to decide what they store  in‐
	side maps.
	-- 一个用户进程可以创建多个maps，然后通过文件描述符访问它们。不同的ebpf程序可以并行访问同一个maps。由用户进程和ebpf程序决定存储内容。

	There's  one  special map type, called a program array.  This type of map stores file de‐
	scriptors referring to other eBPF programs.  When a lookup in the map is  performed,  the
	program flow is redirected in-place to the beginning of another eBPF program and does not
	return back to the calling program.  The level of nesting has a fixed  limit  of  32,  so
	that  infinite loops cannot be crafted.  At run time, the program file descriptors stored
	in the map can be modified, so program functionality can be altered based on specific re‐
	quirements.   All  programs  referred to in a program-array map must have been previously
	loaded into the kernel via bpf().  If a map lookup fails, the current  program  continues
	its execution.  See BPF_MAP_TYPE_PROG_ARRAY below for further details.
	-- 有一类特殊的map类型，称为程序数组。这种类型的map存储指向其它ebpf程序的文件描述符。当在map里执行lookup时，程序流被重定向到另外的ebpf程序的开头，并且不会返回到当前的ebpf程序（调走之后就不回来了）。内嵌的最大层数限制为32，因此不会出现无限循环。在运行时，程序储存在map的文件描述符可以被修改，因此可以根据特定要求更改程序功能。在这种程序数组map里的程序必须在之前已经用bpf()加载了进来。如果lookup失败了，当前程序会继续执行。

	Generally,  eBPF  programs are loaded by the user process and automatically unloaded when
	the process exits.  In some cases, for example, tc-bpf(8), the program will  continue  to
	stay  alive  inside  the kernel even after the process that loaded the program exits.  In
	that case, the tc subsystem holds a reference to the eBPF program after the file descrip‐
	tor  has been closed by the user-space program.  Thus, whether a specific program contin‐
	ues to live inside the kernel depends on how it is further attached  to  a  given  kernel
	subsystem after it was loaded via bpf().
	-- 一般地，ebpf程序被用户进程加载，并随用户进程离开而自动卸载。不过有些特别的例子，比如：tc-bpf，它在加载它的程序退出之后还会继续存活。tc子系统会持有该ebpf程序的引用。

	Each  eBPF program is a set of instructions that is safe to run until its completion.  An
	in-kernel verifier statically determines that the eBPF program terminates and is safe  to
	execute.   During  verification,  the  kernel increments reference counts for each of the
	maps that the eBPF program uses, so that the attached maps can't  be  removed  until  the
	program is unloaded.
	-- 每个ebpf程序都是一个指令集合，安全地执行直到完成。在程序验证期间，内核增加ebpf使用的每个map的引用计数，使得map不会被移除直到程序被卸载。

	eBPF  programs  can  be attached to different events.  These events can be the arrival of
	network packets, tracing events, classification events by network  queueing   disciplines
	(for  eBPF programs attached to a tc(8) classifier), and other types that may be added in
	the future.  A new event triggers execution of the eBPF program, which may store informa‐
	tion  about  the event in eBPF maps.  Beyond storing data, eBPF programs may call a fixed
	set of in-kernel helper functions.
	-- ebpf程序可以附加到不同的事件上。这些事件可以是网络包到达，事件追踪，按网络排队规则对事件进行分类，和其它未来将被加进来的事件。一个新的事件触发ebpf程序的执行，可以存储事件的信息到ebpf maps里。除了存储数据之外，ebpf程序可以调用一系列固定的内核钩子函数。
```

```c
// union 多选一
union bpf_attr {
	struct {    /* Used by BPF_MAP_CREATE */
		__u32         map_type;
		__u32         key_size;    /* size of key in bytes */
		__u32         value_size;  /* size of value in bytes */
		__u32         max_entries; /* maximum number of entries
										in a map */
	};

	struct {    /* Used by BPF_MAP_*_ELEM and BPF_MAP_GET_NEXT_KEY
					commands */
		__u32         map_fd;
		__aligned_u64 key;
		union {
			__aligned_u64 value;
			__aligned_u64 next_key;
		};
		__u64         flags;
	};

	struct {    /* Used by BPF_PROG_LOAD */
		__u32         prog_type;
		__u32         insn_cnt;
		__aligned_u64 insns;      /* 'const struct bpf_insn *' */
		__aligned_u64 license;    /* 'const char *' */
		__u32         log_level;  /* verbosity level of verifier */
		__u32         log_size;   /* size of user buffer */
		__aligned_u64 log_buf;    /* user supplied 'char *'
									buffer */
		__u32         kern_version;
									/* checked when prog_type=kprobe
									(since Linux 4.1) */
	};
} __attribute__((aligned(8)));
```

一个 BPF 程序，如果它通过验证，就会被加载到内核中。在那里，它将被 JIT 编译器**编译成机器码**，并在内核模式下运行，这时附加的触发器将会被激活。

```sh
LOAD -> READ -> WRITE
```

循环缓冲区：内核写入，用户空间程序可以从中读取

BCC, bpftrace 

### Go

> 目前，唯一能够编译成 BPF 机器可以理解的格式的编译器是 Clang。另一种流行的编译器 GСС仍然没有 BPF 后端。而能够编译成 BPF 的编程语言，只有 C 语言的一个非常受限的版本。
> 
> 然而，BPF 程序还有一个在用户空间中的第二部分。这部分可以用 Go 来编写。

BCC 允许你用 Python 编写这一部分，而 Python 是该工具的主要语言。同时，在主库中，BCC 还支持 Lua 和 C++，并且在辅库中，它还支持 Go。

除了 iovisor/gobpf 之外，我还发现了其他三个最新的项目，它们允许你在 Go 中编写用户空间（userland）部分。

- https://github.com/golang/net/tree/master/bpf
- https://github.com/dropbox/goebpf
- https://github.com/cilium/ebpf
- https://github.com/andrewkroh/go-ebpf

### cilium/ebpf

> eBPF is a pure Go library that provides utilities for loading, compiling, and debugging eBPF programs. It has minimal external dependencies and is intended to be used in long running processes.
> 
> eBPF是纯Go库，提供了加载、编译、调试eBPF程序的工具。它只有最小外部依赖，并且适合在长期运行的程序中使用。

尝试执行示例：

1. 统计`sys_execve`系统调用的调用次数

```sh
go run -exec sudo ./kprobe
go: downloading github.com/cilium/ebpf v0.7.1-0.20211126075831-9ead52e53c13
go: downloading golang.org/x/sys v0.0.0-20211001092434-39dca1131b70
[sudo] jd 的密码： 
2021/12/24 14:38:52 Waiting for events..
2021/12/24 14:38:53 sys_execve called 0 times
2021/12/24 14:38:54 sys_execve called 0 times
2021/12/24 14:38:55 sys_execve called 0 times
2021/12/24 14:38:56 sys_execve called 0 times
2021/12/24 14:38:57 sys_execve called 0 times
2021/12/24 14:38:58 sys_execve called 0 times
2021/12/24 14:38:59 sys_execve called 0 times
2021/12/24 14:39:00 sys_execve called 0 times
2021/12/24 14:39:01 sys_execve called 12 times
2021/12/24 14:39:02 sys_execve called 12 times
2021/12/24 14:39:03 sys_execve called 12 times
2021/12/24 14:39:04 sys_execve called 12 times
2021/12/24 14:39:05 sys_execve called 12 times
2021/12/24 14:39:06 sys_execve called 12 times
2021/12/24 14:39:07 sys_execve called 12 times
2021/12/24 14:39:08 sys_execve called 12 times
2021/12/24 14:39:09 sys_execve called 12 times
2021/12/24 14:39:10 sys_execve called 12 times
2021/12/24 14:39:11 sys_execve called 24 times
2021/12/24 14:39:12 sys_execve called 24 times
2021/12/24 14:39:13 sys_execve called 24 times
2021/12/24 14:39:14 sys_execve called 24 times
2021/12/24 14:39:15 sys_execve called 24 times
2021/12/24 14:39:16 sys_execve called 24 times
2021/12/24 14:39:17 sys_execve called 24 times
2021/12/24 14:39:18 sys_execve called 24 times
2021/12/24 14:39:19 sys_execve called 24 times
```

```go
    // 允许当前进程为eBPF资源锁住内存
    rlimit.RemoveMemlock()

    // 加载预编译程序，一般是编译c代码生成的o文件的字节内容
	objs := bpfObjects{}
	loadBpfObjects(&objs, nil)
	defer objs.Close()

	// Open a Kprobe at the entry point of the kernel function and attach the
	// pre-compiled program. Each time the kernel function enters, the program
	// will increment the execution counter by 1. The read loop below polls this
	// map value once per second.
	kp, _ := link.Kprobe(fn, objs.KprobeExecve)
	defer kp.Close()

    // 在执行Kprobe方法时，会将信息写入到objs的KprobeMap里，后续在用户程序即可通过它来查看所需信息
    var value uint64
    objs.KprobeMap.Lookup(mapKey, &value) // 根据mapKey将值读取到value变量
```

## 效果

可以从一个**正在运行**的程序中获得几乎所有的信息，而**无需停止或更改**它。

> eBPF 带来的好处是无与伦比的。
>
>> 首先，从长期看，eBPF 这项新功能会减少未来的 feature creeping normality。 因为用户或开发者希望内核实现的功能，以后**不需要再通过改内核的方式来实现了**。 只需要一段 eBPF 代码，实时动态加载到内核就行了。
>>
>> 其次，因为 eBPF，内核也不会再引入那些影响 fast path 的蹩脚甚至 hardcode 代码 ，从而也避免了性能的下降。
>>
>> 第三，eBPF 还使得内核完全可编程，安全地可编程（fully and safely programmable ），用户编写的 eBPF 程序不会导致内核 crash。另外，eBPF 设计用来解决真实世界 中的线上问题，而且我们现在仍然在坚守这个初衷。

## 更多

[k8s and ebpf](https://arthurchiao.art/blog/ebpf-and-k8s-zh/)

[ebpf](https://cloudnative.to/blog/bpf-intro/)
