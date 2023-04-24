---
author: "jdlau"
date: 2023-04-24
linktitle: NATS
menu:
next:
prev:
title: NATS
weight: 10
categories: ['Go', 'NATS']
tags: ['NATS']
---

## 是什么？

[Home](https://nats.io/about/), [Github](https://github.com/nats-io/nats-server)

> NATS 是一个简单、安全和高性能的通信系统，适用于数字系统、服务和设备。
>
> NATS 是一种允许以消息形式分段的数据交换的基础架构。

### 基于主题

发布者将消息发到主题；订阅者订阅主题，在有消息到来时消费该消息。

> 主题命名规则：
>
> 基本字符：a to z, A to Z and 0 to 9 (区分大小写，不能包含空白字符).
>
> 特殊字符: `.` (分割符，分割不同部分，每部分视为一个token)； * 和 > (通配符，*表示匹配一个token，>表示匹配一或多个token).
>
> 保留主题名称: 以 $ 开头的用在系统内部 (如：$SYS, $JS, $KV ...)

### 发布-订阅

Core NATS: 一个主题，存在一个发布者，多个订阅者。

消息会`复制`到多个订阅者。

### 请求-响应

> A request is sent, and the application either waits on the response with a certain timeout, or receives a response asynchronously.
>
> -- 请求发出后，应用要不等待响应超时，要不就异步收到一个响应。

TODO:

### Queue groups

> if more subscribers are added to the same queue name, they become a queue group, and only one randomly chosen subscriber of the queue group will consume a message each time a message is received by the queue group.
>
> 如果多个订阅者被添加到同一个队列里，它们就成为一个队列组；当消息来时，只会随机选择一个订阅者来消费这条消息。

消息生产后会由任意一个订阅者消费。(`分区`)

## 为什么？

## QOS

> At most once QoS: Core NATS offers an at most once quality of service. If a subscriber is not listening on the subject (no subject match), or is not active when the message is sent, the message is not received. This is the same level of guarantee that TCP/IP provides. Core NATS is a fire-and-forget messaging system. It will only hold messages in memory and will never write messages directly to disk.
> -- 最多一次：如果订阅者没有监听subject，或者当消息被发送时不可用，它将接收不到该消息。这也是TCP/IP提供的同等保证。因为core NATS只会将消息存储在内存里，而不会存储到硬盘里。
>
> At-least / exactly once QoS: If you need higher qualities of service (at least once and exactly once), or functionalities such as persistent streaming, de-coupled flow control, and Key/Value Store, you can use NATS JetStream, which is built in to the NATS server (but needs to be enabled). Of course, you can also always build additional reliability into your client applications yourself with proven and scalable reference designs such as acks and sequence numbers.
> -- 最少一次或刚好一次：使用JetStream模式。它会将消息存储到硬盘里，从而确保消息不会丢失。如果订阅者在消息发送时刚好离线了，在它恢复后，它将会继续消费该条消息。

## 持久化

`jet stream`.

> JetStream was created to solve the problems identified with streaming in technology today - complexity, fragility, and a lack of scalability.
> -- 解决一些流的问题：复杂、碎片、扩展。

## 内嵌NATS服务到应用里

单一进程里启动NATS服务。
