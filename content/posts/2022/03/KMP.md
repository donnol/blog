---
author: "jdlau"
date: 2022-03-22
linktitle: KMP
menu:
next:
prev:
title: KMP
weight: 10
categories: ['algorithm', 'KMP']
tags: ['workspace']
---

KMP字符串匹配算法

精确匹配

状态机

给定一个pattern，查找其在另一字符串s出现的最早位置。（找不到则返回-1）

```go
func index(s string, pattern string) int {

    return -1
}
```

状态推移

```go
func index(s string, pattern string) int {
    n := len(s)
    m := len(pattern)

    // 根据pattern构造dp
    var dp [n][m]int

    // 在s上应用dp，判断pattern位置

    return -1
}
```
