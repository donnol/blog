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

## 更新（2022-11-18）

[临近1.20发布之时，新的提案来了](https://github.com/golang/go/issues/56548)

### 情况

这个提案会兼容泛型加入前已有的规则。也就是能用`==`，但是有可能在运行时panic。

```go
type M struct{ f any }                        // 注意看，f字段的类型为any
fmt.Println(M{f: 1} == M{f: 2})               // 可编译通过，运行时正常执行：false
fmt.Println(M{f: []int{1}} == M{f: []int{2}}) // 可编译通过，运行时panic：panic: runtime error: comparing uncomparable type []int
```

[跑来看看](https://go.dev/play/p/69Un_xqq0cX)

> If we want any to satisfy comparable, then constraint satisfaction can't be the same as interface implementation. A non-comparable type T (say []int) implements any, but T does not implement comparable (T is not in comparable's type set). Therefore any cannot possibly implement comparable (the implements relation is transitive - we cannot change that). So if we want any to satisfy the comparable constraint, then constraint satisfaction can't be the same as interface implementation.
>
> -- 如果想要any满足comparable，约束满足就不能跟接口实现一样。
> -- 这里用[]int类型举例，它是不可比较的（其实它可以与nil比较，不过也仅能与nil比较），它实现了any，但它没有实现comparable。所以按照这样下去，any是不能实现comparable的。
> -- 那么，如果我们想要any满足comparable约束，约束满足就不能跟接口实现一样。

### 提议

> We change the spec to use a different rule for constraint satisfaction than for interface implementation: we want spec-comparable types to satisfy comparable; i.e., we want to be able to use any type for which == is defined even if it may not be strictly comparable.
>
> -- 修改spec规范，对于`约束满足`使用跟`接口实现`不同的规则。约定`spec-comparable`类型满足`comparable`。
>
> With this change, constraint satisfaction matches interface implementation but also contains an exception for spec-comparable types. This exception permits the use of interfaces as type arguments which require strict comparability.
>
> -- 这样修改之后，`约束满足`除了会有一个关于`spec-comparable`类型的异常外，基本上匹配`接口实现`。这个异常允许使用接口作为泛型的类型参数，这个类型参数要求`strict-comparability(严格的可比较)`。

关于`spec-comparable`和`strict-comparable`:

> For clarity, in the following we use the term `strictly comparable` for the types in comparable, and spec-comparable for types of `comparable operands`. Strictly comparable types are spec-comparable, but not the other way around. Types that are not spec-comparable are simply not comparable.
>
> -- 在`comparable`里的类型即是`strictly comparable`的，支持比较符（==，!=）的类型是`spec-comparable`。
> 很明显，`Strictly comparable`类型一定是`spec-comparable`，但反过来就不一定。不是`spec-comparable`的类型就一定不是`comparable`(不能满足comparable约束)。

关于`satisfy`:

> We also add a new paragraph defining what "satisfy" means:
>
>> A type T satisfies a constraint interface C if
>>
>> T implements C; or
>> C can be written in the form interface{ comparable; E }, where E is a basic interface and T is comparable and implements E.
>
> -- 如果说类型T满足约束C：
> -- T实现了C；或者
> -- C可以被写为以下格式：interface{ comparable: E}，其中`E是一个基础接口(只有方法，没有type set)`并且`T满足comparable约束并实现了E`。

### 编译时静态检查

**那如何在编译时确保某类型是可比较的呢？**

如果要在编译时确保某个类型是可比较的，可以[这样](https://github.com/golang/go/issues/56548#issuecomment-1317673963)：

```go
// we want to ensure that T is strictly comparable
type T struct {
	x int
}

// define a helper function with a type parameter P constrained by T
// and use that type parameter with isComparable
// -- 把该类型T作为约束使用，并且对应的类型参数用于实例化一个使用了comparable约束的泛型函数
func TisComparable[P T]() {
	_ = isComparable[P]
}

func isComparable[_ comparable]() {}
```
