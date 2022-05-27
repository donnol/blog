---
author: "jdlau"
date: 2022-05-27
linktitle: Find out which Go version built your binary
menu:
next:
prev:
title: Find out which Go version built your binary
weight: 10
categories: ['Go']
tags: ['version']
---

# 根据二进制文件找出应用构建时使用的`Go`版本

使用[`dlv`](https://github.com/go-delve/delve/blob/master/Documentation/installation):

```sh
dlv exec ./app
> p runtime.buildVerion
```

或者，在代码里调用`runtime.Version()`:

```go
func main() {
    fmt.Println("go version:", runtime.Version())
}
```

[参照](https://dave.cheney.net/2017/06/20/how-to-find-out-which-go-version-built-your-binary)
