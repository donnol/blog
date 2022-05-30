---
author: "jdlau"
date: 2022-05-30
linktitle: 泛型
menu:
next:
prev:
title: 泛型
weight: 10
categories: ['Go', 'TypeScript']
tags: ['generic']
---

# Go泛型

试看：

```go
// 准确的描述出了参数和返回值的类型，非常方便
func Add(x, y int) int

// 但也限制了Add函数的参数类型--只能接收`int`
// Add(0.1, 0.2) // can't do that

// 那怎么办呢？再写一个针对float64的呗
func AddFloat64(x, y float64) float64

AddFloat64(0.1, 0.2) // it's ok

// 如果还要支持其它类型呢？再加一个吗，每多一种类型，就多加一个。。。
func AddInt8(x, y int8) int8
func AddInt32(x, y int32) int32
func AddFloat32(x, y float32) float32
// more...

// emm.

// how about interface{}?
func AddAny(x, y interface{}) interface{} {
    switch x.(type) {
        case int:
        case int8:
        case int32:
        case float32:
        case float64:
        // more...
        default:
        panic("not support type")
    }
}
// interface{}表示可以接收任意类型的值，并且返回任意类型的值
// 换言之，参数的类型和返回值的类型没有必然联系--从签名看来，它们可以一样，也可以不一样

func AddGeneric1[T any](x, y T) T // 看起来跟AddAny差不多，但是参数类型和返回值类型必然是相同的

// 但any并不一定支持+运算符，所以需要用更细粒度的约束
type Number interface {
    ~int|~int8|~int32|~float32|~float64
}
func AddGeneric2[T Number](x, y T) T // 通过Number约束，确保类型参数可加

// 泛型的存在，使得函数的类型集比any小，比int大；使得返回值和参数的类型能够动态联系。

func Map[T, E any](list []T, f func(T) E) []E {
    r := make([]E, len(list))
    for i := range list {
        r[i] = f(list[i])
    }
    return r
}
```

## ts的泛型

keyof的存在，ts的泛型更为强大：

```ts
// Cond出现的字段必须分别在L和R均出现
export type Cond<L, R> = {
    l: keyof L
    r: keyof R
}

export function leftJoin<L, R, LR>(
    left: L[],
    right: R[],
    cond: Cond<L, R>,
    f: (l: L, r: R) => LR,
): LR[] {
    let rm = new Map()
    right.forEach((value: R, index: number, array: R[]) => {
        let rv = getProperty(value, cond.r)
        rm.set(rv, value)
    })

    let res = new Array(left.length)
    left.forEach((value: L, index: number, array: L[]) => {
        let lv = getProperty(value, cond.l)
        let rv = rm.get(lv)
        let lr = f(value, rv)
        res.push(lr)
    })

    return res
}

export function getProperty<T, K extends keyof T>(o: T, propertyName: K): T[K] {
    return o[propertyName]; // o[propertyName] is of type T[K]
}

const users = [
    {
        id: 1,
        name: "jd",
    },
    {
        id: 2,
        name: "jk",
    },
]

const articles = [
    {
        id: 1,
        userId: 1,
        name: "join"
    },
    {
        id: 2,
        userId: 2,
        name: "join 2"
    },
]

let res = leftJoin(users, articles, { l: "id", r: "userId" }, (l, r) => {
    return {
        id: r.id,
        name: r.name,
        userId: l.id,
        userName: l.name,
    }
})
console.log("res: ", res)
```
