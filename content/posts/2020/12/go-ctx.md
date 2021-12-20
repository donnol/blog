---
author: "jdlau"
date: 2020-12-18
linktitle: go ctx
menu:
next: 
prev: 
title: go ctx
weight: 10
categories: ['go']
tags: ['ctx']
---

## ctx

1.why

`goroutine`号称百万之众，互相之间盘根错节，难以管理控制。为此，必须提供一种机制来管理控制它们。

### 各自为战

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // start first
    go func() {
        fmt.Println(1)
    }()

    // start second
    go func() {
        fmt.Println(2)
    }()

    time.Sleep(time.Second)
}
```

### 万法归一

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := new(sync.WaitGroup)

    // start first
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(1)
    }()

    // start second
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(2)
    }()

    wg.Wait()
}

```

可以看到使用`waitgroup`可以控制多个`goroutine`必须互相等待，直到最后一个完成才会全部完成。

### 明修栈道暗度陈仓

```go
package main

import (
    "fmt"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    // start first
    go func() {
        fmt.Println(1)
        <-ch2
        ch1 <- 1
    }()

    ch3 := make(chan int)

    // start second
    go func() {
        fmt.Println(2)
        ch2 <- 2
        <-ch1

        // escape
        ch3 <- 3
    }()

    n := <-ch3
    fmt.Println(n)
}
```

使用`chan`的话，可以实现`goroutine`之间的消息同步

2.what

[Package context](https://pkg.go.dev/context) defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

-- 提供标准库`context`，定义了`Context`类型，带有限期、取消信息和其它请求域里的跨API边界和进程间的值。

3.how

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    var n time.Duration = 2

    now := time.Now()
    ctx, cancel := context.WithDeadline(context.Background(), now.Add(time.Second*n))
    _ = cancel

    fmt.Println(now)

    // start first
    go func(ctx context.Context) {
        select {
        case <-ctx.Done():
        }
        fmt.Println(time.Now(), 1)
    }(ctx)

    // start second
    go func(ctx context.Context) {
        select {
        case <-ctx.Done():
        }
        fmt.Println(time.Now(), 2)
    }(ctx)

    time.Sleep(time.Second * (n - 1))
    fmt.Println(time.Now())
 
    // 一秒钟之后取消的话，两个goroutine会在取消后马上执行；如果等到时间到期了，就会在两秒后执行；
    // cancel()
    // fmt.Println(time.Now())

    time.Sleep(time.Second * (n + 1))
}
```

4.others
