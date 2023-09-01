---
author: "jdlau"
date: 2023-09-01
linktitle: Go Empty Struct
menu:
next:
prev:
title: Go Empty Struct
weight: 10
categories: ['Go']
tags: ['mark']
---

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type A struct{}
	type B struct{}
    // 结构体里的字段都是`Empty Struct`时，占用空间为0
	type S struct {
		A A
		B B
	}
	var s S
	fmt.Println(unsafe.Sizeof(s)) // prints 0
    // 如果是指针，占用空间为8
	fmt.Println(unsafe.Sizeof(&s)) // prints 8

	var x [1000000000]struct{}
    // 可以同时存储A和B类型元素
	x[0] = A{}
	x[1] = B{}
	fmt.Println(unsafe.Sizeof(x)) // prints 0
    // 地址一样
	fmt.Printf("%p, %p", &x[0], &x[1]) // 0x54e3a0, 0x54e3a0
}
```

[See also](https://dave.cheney.net/2014/03/25/the-empty-struct)
