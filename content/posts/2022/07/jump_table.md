---
author: "jdlau"
date: 2022-07-20
linktitle: Jump Table
menu:
next:
prev:
title: Jump Table
weight: 10
categories: ['Go']
tags: ['version']
---

## What's Jump Table?

A jump table can be either `an array of pointers to functions` or `an array of machine code jump instructions`. If you have **a relatively static set of functions** (such as system calls or virtual functions for a class) then you can create this table once and call the functions using a simple index into the array. This would mean retrieving the pointer and calling a function or jumping to the machine code depending on the type of table used.

The benefits of doing this in embedded programming are:

- Indexes are more memory efficient than machine code or pointers, so there is a potential for memory savings in constrained environments.
- For any particular function the index will remain stable and changing the function merely requires swapping out the function pointer.
If does cost you a tiny bit of performance for accessing the table, but this is no worse than any other virtual function call.

[From Stackoverflow](https://stackoverflow.com/questions/48017/what-is-a-jump-table)

## Try

```go
// 跳表
// 创建后，不会改变的函数指针数组；后续可以根据数组索引来直接找到相应函数并执行
// 至于，为什么要这样干，是因为数组索引查找更高效

// 初始化
var (
	funcArray = [...]func(){
		func() {
			fmt.Println(0)
		},
		func() {
			fmt.Println(1)
		},

		// more ...
	}
)

func main() {
	// 使用
	funcArray[0]()
	funcArray[1]()

	// more ...
}
```

[Playground](https://go.dev/play/p/oJEo1iaRfSW)
