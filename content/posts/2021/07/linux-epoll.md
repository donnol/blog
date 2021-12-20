---
author: "jdlau"
date: 2021-07-17
linktitle: linux epoll
menu:
next: 
prev: 
title: linux epoll
weight: 10
categories: ['linux']
tags: ['epoll']
---

# linux epoll

[wiki](https://zh.wikipedia.org/wiki/Epoll)

[手册](https://man7.org/linux/man-pages/man7/epoll.7.html)

## why

## what

Linux内核的**可扩展I/O事件通知机制**。

于Linux 2.5.44首度登场，它设计目的旨在取代既有POSIX select(2)与poll(2)系统函数，让需要大量操作文件描述符的程序得以发挥更优异的性能（举例来说：旧有的系统函数所花费的时间复杂度为O(n)，epoll的时间复杂度O(log n)）。epoll 实现的功能与 poll 类似，都是监听多个文件描述符上的事件。

## how

epoll 通过使用红黑树(RB-tree)搜索被监控的文件描述符(file descriptor)。

在 epoll 实例上注册事件时，epoll 会将该事件添加到 epoll 实例的红黑树上并注册一个回调函数，当事件发生时会将事件添加到就绪链表中。

```c
int epoll_create(int size);
```

在内核中创建epoll实例并返回一个epoll文件描述符。

```c
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
```

向 epfd 对应的内核epoll 实例添加、修改或删除对 fd 上事件 event 的监听。op 可以为 EPOLL_CTL_ADD, EPOLL_CTL_MOD, EPOLL_CTL_DEL 分别对应的是添加新的事件，修改文件描述符上监听的事件类型，从实例上删除一个事件。如果 event 的 events 属性设置了 EPOLLET flag，那么监听该事件的方式是边缘触发。

```c
int epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout);
```

当 timeout 为 0 时，epoll_wait 永远会立即返回。而 timeout 为 -1 时，epoll_wait 会一直阻塞直到任一已注册的事件变为就绪。当 timeout 为一正整数时，epoll 会阻塞直到计时 timeout 毫秒终了或已注册的事件变为就绪。因为内核调度延迟，阻塞的时间可能会略微超过 timeout 毫秒。

### 触发模式

epoll提供**边沿触发**及**状态触发**模式。

在边沿触发模式中，**epoll_wait仅会在新的事件首次被加入epoll队列时返回**；在状态触发模式下，epoll_wait在**事件状态未变更前**将不断被触发。状态触发模式是默认的模式。

状态触发模式与边沿触发模式有**读和写**两种情况，我们先来考虑读的情况。假设我们注册了一个读事件到epoll实例上，epoll实例会**通过epoll_wait返回值的形式通知我们哪些读事件已经就绪**。简单地来说，在状态触发模式下，如果读事件未被处理，该事件对应的内核读缓冲器非空，则**每次调用`epoll_wait`时返回的事件列表都会包含该事件**，直到该事件对应的内核读缓冲器为空为止。而在边沿触发模式下，**读事件就绪后只会通知一次**，不会反复通知。

然后我们再考虑写的情况。状态触发模式下，只要文件描述符对应的内核写缓冲器未满，就会**一直通知可写事件**。而在边沿触发模式下，内核写缓冲器由满变为未满后，只会**通知一次可写事件**。

举例来说，倘若有一个已经于epoll注册之流水线接获资料，epoll_wait将返回，并发出资料读取的信号。现假设缓冲器的资料仅有部分被读取并处理，在level-triggered(状态触发)模式下，任何对epoll_wait之调用都将**即刻返回**，直到缓冲器中的资料全部被读取；然而，在edge-triggered(边缘触发)的情境下，epoll_wait**仅会于再次接收到新资料**(亦即，新资料被写入流水线)时返回。

#### 边沿触发模式

边沿触发模式使得**程序有可能在用户态缓存 IO 状态**。nginx 使用的是边沿触发模式。

文件描述符有两种情况是推荐使用边沿触发模式的。

1. read 或者 write 系统调用返回了 EAGAIN。
2. 非阻塞的文件描述符。

可能的缺陷：

如果 IO 空间很大，你要花很多时间才能把它一次读完，这可能会导致**饥饿**。举个例子，假设你在监听一个文件描述符列表，而**某个文件描述符上有大量的输入**（不间断的输入流），那么你**在读完它的过程中就没空处理其他就绪的文件描述符**。（因为边沿触发模式**只会通知一次可读事件**，所以你**往往会想一次把它读完**。）一种解决方案是，程序维护一个就绪队列，当 epoll 实例通知某文件描述符就绪时将它**在就绪队列数据结构中标记为就绪**，这样程序就会记得哪些文件描述符等待处理。Round-Robin 循环处理就绪队列中就绪的文件描述符即可。

如果你缓存了所有事件，那么一种可能的情况是 A 事件的发生让程序关闭了另一个文件描述符 B。但是内核的 epoll 实例并不知道这件事，需要你从 epoll 删除掉。
