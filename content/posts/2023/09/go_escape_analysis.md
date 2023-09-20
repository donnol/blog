---
author: "jdlau"
date: 2023-09-20
linktitle: Go escape analysis
menu:
next:
prev:
title: Go escape analysis
weight: 10
categories: ['Go', 'escape analysis']
tags: ['note']
---

> The meaning of `escapes to the heap` is variables needs to `be shared across the function stack frames` [between main() and Println()]
> ...
>
> ...
> So `globally access variables` must be `moved to heap` as it requires runtime. So the output line 11:2 shows the same as the data variable moved to the heap memory.

[From](https://mayurwadekar2.medium.com/escape-analysis-in-golang-ee40a1c064c1)
