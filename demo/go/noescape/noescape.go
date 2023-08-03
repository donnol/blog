package noescape

import "unsafe"

// Code is from strings.Builder
//
// go build -gcflags='-m -l' noescape/noescape.go
// noescape/noescape.go:13:15: p does not escape
// noescape/noescape.go:24:13: leaking param: p to result ~r0 level=0

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	// 转换一下，就能确保不会逃逸到堆
	x := uintptr(p)

	// x ^ 0 永远等于x，那直接用x呢？测试后发现也不会逃逸啊，为什么要用x ^ 0呢？
	return unsafe.Pointer(x ^ 0) //nolint:analysis
}

func escape(p unsafe.Pointer) unsafe.Pointer {
	return p
}
