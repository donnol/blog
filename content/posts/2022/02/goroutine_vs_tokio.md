---
author: "jdlau"
date: 2022-02-16
linktitle: goroutine vs tokio
menu:
next:
prev:
title: goroutine vs tokio
weight: 10
categories: ['go', 'goroutine', 'tokio']
tags: ['concurrent']
---

[Reddit讨论贴](https://www.reddit.com/r/rust/comments/lg0a7b/benchmarking_tokio_tasks_and_goroutines/)

> Go uses a different strategy for blocking systemcalls. It does not run them on a threadpool - it moves all the other goroutines that are queued to run on the current thread to a new worker thread, then runs the blocking systemcall on the current thread. **This minimizes context switching**.
>
> You can do this in tokio as well, using **task::block_in_place**. If I change your code to use that instead of tokio::fs, it gets a lot closer to the go numbers. Note that using block_in_place is not without caveats, and it only works on the multi-threaded runtime, not the single-threaded one. That's why it's not used in the implementation of tokio::fs.
>
> On my Linux desktop:
>>
>> goroutines: `3.22234675s total`, 3.222346ms avg per iteration
>>
>> rust_threads: 16.980509645s total, 16.980509ms avg per iteration
>>
>> rust_tokio: 9.56997204s total, 9.569972ms avg per iteration
>>
>> rust_tokio_block_in_place: `3.578928749s total,` 3.578928ms avg per iteration

[对比文章](https://qvault.io/rust/concurrency-in-rust-can-it-stack-up-against-gos-goroutines/)

> Goroutines are more lightweight and efficient than operating-system threads. As a result, a program can **spawn more total goroutines** than threads. Goroutines also start and clean themselves up faster than threads due to less system overhead.
>
> The big advantage of traditional threading (like that of Rust) over the goroutine model is that **no runtime is required**. Each Go executable is compiled with a small runtime which manages goroutines, while Rust avoids that extra fluff in the binary.

[附：实例对比](https://rustcc.cn/article?id=5985dfe8-e6f8-46c6-9172-3e05d3ee91f7)

[io比较](https://medium.com/star-gazers/benchmarking-low-level-i-o-c-c-rust-golang-java-python-9a0d505f85f7)

[让tokio调度更快](https://tokio.rs/blog/2019-10-scheduler)

> The run queue must support both multiple producers and multiple consumers. The commonly used algorithm is an **intrusive linked list**(侵入性的链表). Intrusive implies that the task structure includes a pointer to the next task in the run queue instead of wrapping the task with a linked list node. This way, allocations are avoided for push and pop operations. It is possible to use a lock-free push operation but popping requires^1 a mutex to coordinate consumers.
>
> This scheduler model has a downside. **All processors contend on the head of the queue**(contend: 竞争). For general-purpose thread pools, this usually is not a deal breaker. The amount of time processors spend executing the task far outweighs the amount of time spent popping the task from the run queue. When tasks execute for a long period of time, queue contention is reduced. However, Rust's asynchronous tasks are expected to take very little time executing when popped from the run queue. In this scenario, the overhead from contending on the queue becomes significant.

[rust async await](https://liufuyang.github.io/2019/11/10/manish-async-translation.html)

> 在函数调用时插入调度点，rust通过yield来插入，实现一个函数多次返回。
