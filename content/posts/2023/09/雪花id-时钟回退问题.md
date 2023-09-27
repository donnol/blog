---
author: "jdlau"
date: 2023-09-27
linktitle: 雪花id的时钟回退问题
menu:
next:
prev:
title: 雪花id的时钟回退问题
weight: 10
categories: ['Snowflake', 'Time']
tags: ['mark']
---

雪花id由64位二进制组成，转成字符串则长为19. 它依赖于系统时钟，如果出现时钟回退，会导致已经在用的id再次被生成。

怎么办呢？

0. 记录上次生成时间，在本次生成时比较时间，如果当前时间比上次生成时间要小，则认为时钟回拨，直接报错。也可以一直重试，直到当前时间不小于上次生成时间。

1. 采用历史时间则天然的不存在时间回拨问题。但是在超高并发情况下，历史的时间很快用完，时间一直保持在最新时间的话，这个时候还是会出现时间回拨。

2. Go1.9开始，使用单调时钟: time.Now(), time.Since(), time.Until().

> // # Monotonic Clocks
>
> //
>
> // Operating systems provide both a “wall clock,” which is subject to
>
> // changes for clock synchronization, and a “monotonic clock,” which is
>
> // not. The general rule is that the wall clock is for telling time and
>
> // the monotonic clock is for measuring time. Rather than split the API,
>
> // in this package the Time returned by time.Now contains both a wall
>
> // clock reading and a monotonic clock reading; later time-telling
>
> // operations use the wall clock reading, but later time-measuring
>
> // operations, specifically comparisons and subtractions, use the
>
> // monotonic clock reading.
>
> //
>
> // For example, this code always computes a positive elapsed time of
>
> // approximately 20 milliseconds, `even if the wall clock is changed` during
>
> // the operation being timed:
>
> //
>
> //	start := time.Now()
>
> //	... operation that takes 20 milliseconds ...
>
> //	t := time.Now()
>
> //	elapsed := t.Sub(start)
>
> //
>
> // Other idioms, such as time.Since(start), time.Until(deadline), and
>
> // time.Now().Before(deadline), are similarly robust against wall clock
>
> // resets.
>

[R1](https://go.dev/src/time/time.go?s=17583:17610)

[R2](https://github.com/bwmarrin/snowflake/pull/18/files)
