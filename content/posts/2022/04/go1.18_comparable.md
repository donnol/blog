---
author: "jdlau"
date: 2022-04-22
linktitle: Go1.18 comparable
menu:
next:
prev:
title: Go1.18 comparable
weight: 10
categories: ['Go']
tags: ['comparable']
---

# Go 1.18 预定义接口类型

先看一个提案: [proposal: spec: permit values to have type "comparable"](https://github.com/golang/go/issues/51338) -- 允许值拥有`comparable`类型，我的理解是，现在的`comparable`只能用作泛型里的类型参数的约束，不能像普通类型那样使用，如下：

```go
type Set[E comparable] []E // 可以用做类型参数的约束

// 使用go1.18编译，报错：interface is (or embeds) comparable
var A comparable // 变量不可以使用`comparable`类型
```

那么，结合例子就能更好地理解这个提案了。

这个提案的主要目的就是让例子里的`var A comparable`成立，也就是允许`comparable`作为变量的类型，跟其它普通的接口类型(`var E error`)一样。

```go
// proposal: spec: permit values to have type "comparable"

// As part of adding generics, Go 1.18 introduces a new predeclared interface type comparable. That interface type is implemented by any non-interface type that is comparable, and by any interface type that is or embeds comparable. Comparable non-interface types are numeric, string, boolean, pointer, and channel types, structs all of whose field types are comparable, and arrays whose element types are comparable. Slice, map, and function types are not comparable.
// -- 作为泛型的一部分，Go1.18引入了一个新的预定义接口类型：`comparable`。这个接口类型由任何可比较的非接口类型实现，和任何是`comparable`或内嵌了`comparable`的接口类型实现。可比较的非接口类型有：数值、字符串、布尔、指针、字段类型均是可比较的管道或结构体类型、元素是可比较的数组类型。切片、映射、函数类型均是不可比较的。

// In Go, interface types are comparable in the sense that they can be compared with the == and != operators. However, interface types do not in general implement the predeclared interface type comparable. An interface type only implements comparable if it is or embeds comparable.
// 在Go里面，接口类型是可比较的意味着它们可以用`==`和`!=`操作符进行比较。但是，接口类型一般来说没有实现预定义接口类型`comparable`。一个接口类型只有它是或内嵌了`comparable`时才实现了`comparable`。

// Developing this distinction between the predeclared type comparable and the general language notion of comparable has been confusing; see #50646. The distinction makes it hard to write certain kinds of generic code; see #51257.
// 出现的两个问题：[怎么在文档里说明哪些接口实现它了呢？](https://github.com/golang/go/issues/50646), [any是任意类型的意思，那必然比comparable大吧](https://github.com/golang/go/issues/51257)
// 突然想到：如果我要把一个变量表示为不可比较的，怎么样可以用`comparable`来表示呢，`!comparable`?

// For a specific example, you can today write a generic Set type of some specific (comparable) element type and write functions that work on sets of any element type:
// 
// type Set[E comparable] map[E]bool
// func Union[E comparable](s1, s2 Set[E]) Set[E] { ... }
// 
// But there is no way today to instantiate this Set type to create a general set that works for any (comparable) value. That is, you can't write Set[any], because any does not satisfy the constraint comparable. You can get a very similar effect by writing map[any]bool, but then all the functions like Union have to be written anew for this new version.

// We can reduce this kind of problem by permitting comparable to be an ordinary type. It then becomes possible to write Set[comparable].

// As an ordinary type, comparable would be an interface type that is implemented by any comparable type.
// 作为一个普通类型，`comparable`是一个可以被任意`comparable`类型实现的接口类型。

// Any comparable non-interface type could be assigned to a variable of type comparable.
// -- 任何可比较的非接口类型可以被分配到类型为`comparable`的变量。
// A value of an interface type that is or embeds comparable could be assigned to a variable of type comparable.
// -- 接口类型是或内嵌了`comparable`的值可以被分配到类型为`comparable`的变量。
// A type assertion to comparable, as in x.(comparable), would succeed if the dynamic type of x is a comparable type.
// 类型断言，如`x.(comparable)`，当x的动态类型是一个`comparable`类型时可以成功。
// Similarly for a type switch case comparable.
// 对`type switch`来说类似。
```

```go
type C interface {
    comparable
}

var c C

func main() {
    var A comparable

    var a int

    if v, ok := a.(comparable); ok {

    }

    switch a.(type) {
    case comparable:

    }
}
```

## 反射的`Comparable`

```go
func ReflectComparable(v interface{}) bool {
	typ := reflect.TypeOf(v)

	// Comparable reports whether values of this type are comparable.
	// Even if Comparable returns true, the comparison may still panic.
	// For example, values of interface type are comparable,
	// but the comparison will panic if their dynamic type is not comparable.
	// -- 即使返回true，也有可能panic。
	// 比如：接口类型的值是可比较的，但如果它们的动态类型是不可比较的，就会panic
	return typ.Comparable()
}
```

## go/types的`Comparable`

```go
func TypesComparable() bool {
	t := types.NewChan(types.SendOnly, &types.Basic{})

	return types.Comparable(t)
}
```

## 更新

[使comparable仅在类型集里没有任何一个不可比较类型时正确，否则依然在编译时可通过，但运行时panic](https://github.com/golang/go/issues/52614)

```go
func Comparable[T comparable](t1, t2 T) bool {
	return t1 == t2
}

type IC interface {
	Name() string
}

type ComparableStruct struct {
	name string
}

func (cs ComparableStruct) Name() string {
	return cs.name
}

type NotComparableStruct struct {
	name string

	m map[int]string
}

func (ncs NotComparableStruct) Name() string {
	return ncs.name
}

func main() {
	var a IC = ComparableStruct{name: "jd"}
	var b IC = NotComparableStruct{name: "jd", m: make(map[int]string)}
	_, _ = a, b
	// 现在提案没实现，所以普通接口并未实现comparable，会编译报错: IC does not implement comparable
	// 如果提案通过，将会是编译通过，执行panic
	// Comparable(a, b)
}
```
