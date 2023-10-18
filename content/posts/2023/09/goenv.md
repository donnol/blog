---
author: "jdlau"
date: 2023-09-13
linktitle: Vscode go cannot find GOROOT directory
menu:
next:
prev:
title: Vscode go cannot find GOROOT directory
weight: 10
categories: ['Go', 'Vscode']
tags: ['mark']
---

今天发现在windows上的vscode一直提示找不到go：`go: cannot find GOROOT directory: c:\msys64\mingw64\lib\go`。

强制设置了go.goroot也不行，直到查看了GOENV文件（C:\Users\xxx\AppData\Roaming\go\env）之后，才发现里面有一行：GOROOT=c:\msys64\mingw64\lib\go，可能是当时在msys2安装go的时候加上的。

去掉它就恢复正常了。

```sh
$ go env
set GOENV=C:\Users\xxx\AppData\Roaming\go\env
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOMODCACHE=C:\Users\xxx\go\pkg\mod
set GOOS=windows
set GOPATH=C:\Users\xxx\go
set GOPRIVATE=
set GOPROXY=https://goproxy.cn,https://goproxy.io,direct
set GOROOT=C:\Program Files\Go
```

应该是这样的，如果用`go env -w `来设置goroot，那么这个值就会保存到GOENV对应的文件里，如果是`$env:GOROOT=xxx`的方式来设置则不会修改GOENV文件里的内容。这时候，如果vscode是优先从GOENV文件来获取GOROOT的话，就可能会导致与实际的GOROOT不一致。

所以，如果再遇到以上错误，除了`echo $env:GOROOT` 看一下环境变量值之外，也要看一下`GOENV`文件。
