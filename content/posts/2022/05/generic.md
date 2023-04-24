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

# 泛型

## 是什么？

Type parameter, 类型参数。`func Add[T Number](x, y T) (r T)`，其中的`T`就是类型参数，它被接口`Number`所约束。

```go
type Number interface {
    int | float32
}
```

调用方除了可自行决定参数值之外，还可以自行决定参数类型。`Add[int](1, 2)`，在调用时指定`T`的类型为`int`，同时传入参数值`1`,`2`必须是`int`类型。

这样使得代码更灵活，更有扩展性，同时更安全。

## Go泛型

### 为什么？

静态语言，类型固定，比如这个函数：`func Add(x, y int) int`就要求参数和结果都必须是整型。

那如果后来又需要一个浮点数的加法呢？

![固定类型](/image/固定类型.png)

那使用interface{}不也可以吗？

![固定类型](/image/any类型.png)

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
// 所以，使用interface{}不够安全。

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

![类型参数](/image/类型参数.png)

### 一点不足

泛型虽然出来了，但是类型推断依然不够强大，1.21有望做出[改进](https://github.com/golang/go/issues/58650)和[提升](https://github.com/golang/go/issues/59338)。

```go
// 第三个参数initial其实在函数实现里并没使用到，但是为了在调用时可以省略类型参数，所以需要这样一个参数
// 但其实，如果类型推断足够强大的话，是可以从Finder约束的NewScanObjAndFields方法推断出R类型的
func FindAll[S Storer, F Finder[R], R any](db S, finder F, initial R) (r []R, err error) {
    query, args := finder.Query()
	rows, err := db.QueryContext(context.TODO(), query, args...) // sql里select了n列
	if err != nil {
		return
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	for rows.Next() {
		obj, fields := finder.NewScanObjAndFields(colTypes) // fields也必须有n个元素
		if err = rows.Scan(fields...); err != nil {
			return
		}
		// PrintFields(fields)

		r = append(r, *obj)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

// 调用时需要该无用参数
r, err := do.FindAll(tdb, finder, (UserForDB{}))

// 这个才是我们想要的函数签名
func FindAll1[S do.Storer, F do.Finder[R], R any](db S, finder F) (r []R, err error) 

// 但是现在无法推断出R的类型，除非显式标明类型参数
FindAll1(tdb, finder1) // Error: cannot infer R (/home/jd/Project/jd/tools/db/find.go:37:38)
// 显然，如果把类型参数写出来，是非常啰嗦的
FindAll1[*sql.DB, *finderOfUser, UserForDB](tdb, &finderOfUser{})

type Finder[R any] interface {
	Query() (query string, args []any)

	NewScanObjAndFields(colTypes []*sql.ColumnType) (r *R, fields []any)
}

type Storer interface {
	*sql.DB | *sql.Tx | *sql.Conn
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}
```

### 应用

hash join:

```go
func HashJoin[K comparable, LE, RE, R any](
    // 左表和右表
	left []LE,
	right []RE,
    // 左关联函数和右关联函数，它们返回同类型键
	lk func(item LE) K,
	rk func(item RE) K,
    // 接收左表元素和右表元素，返回新元素
	mapper func(LE, RE) R,
) []R {
	var r = make([]R, 0, len(left))

    // 先对右表做映射
	rm := KeyBy(right, rk)

	for _, le := range left {
		k := lk(le)
		re := rm[k]
        // 将关联元素做处理
		r = append(r, mapper(le, re))
	}

	return r
}

// KeyValueBy slice to map, key value specified by iteratee
func KeyValueBy[K comparable, E, V any](collection []E, iteratee func(item E) (K, V)) map[K]V {
	result := make(map[K]V, len(collection))

	for i := range collection {
		k, r := iteratee(collection[i])
		result[k] = r
	}

	return result
}

// KeyBy slice to map, key specified by iteratee, value is slice element
func KeyBy[K comparable, E any](collection []E, iteratee func(item E) K) map[K]E {
	return KeyValueBy(collection, func(item E) (K, E) {
		return iteratee(item), item
	})
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
