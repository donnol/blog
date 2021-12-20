---
author: "jdlau"
date: 2021-12-14
linktitle: mqtt
menu:
next: /posts/2021/12/redis_sds/
prev: /posts/2021/12/k8s/
title: mqtt
weight: 10
categories: ['message queue']
tags: ['mqtt']
---

## 物联网消息标准

[官网](https://mqtt.org/)

> It is designed as an extremely lightweight publish/subscribe messaging transport that is ideal for connecting remote devices with a small code footprint and minimal network bandwidth.
>
> 极其轻量的发布/订阅消息传输，使用小量代码脚本和极小网络带宽来连接远程设备。

- 轻量
- 高效
- 双向
- 大规模（百万设备）
- 可靠
- 支持不可靠网络
- 安全

![架构图](/image/mqtt架构图.png)

多个`mqtt`客户端连接到`broker`(译为：中间商)，围绕`topic`来实现发布/订阅操作，某些客户端向`topic`发布消息，某些客户端订阅`topic`上的消息，当`broker`接收到某个`topic`上的消息时，它会将消息转发到订阅了该`topic`的客户端。

[mqtt 5.0](https://docs.oasis-open.org/mqtt/mqtt/v5.0/mqtt-v5.0.html)

## QoS

Quality of Service

> control traffic and ensure the performance of critical applications with limited network capacity
>
> 控制交通，确保有限网络容量下的应用性能。

QoS（Quality of Service，服务质量）指一个网络能够利用各种基础技术，为指定的网络通信提供更好的服务能力，是网络的一种安全机制， 是用来**解决网络延迟和阻塞**等问题的一种技术。QoS 的保证对于容量有限的网络来说是十分重要的，特别是对于流多媒体应用，例如 VoIP 和 IPTV 等，因为这些应用常常需要固定的传输率，对延时也比较敏感。

> 当网络发生拥塞的时候，所有的数据流都有可能被丢弃；为满足用户对不同应用不同服务质量的要求，就需要**网络能根据用户的要求分配和调度资源**，**对不同的数据流提供不同的服务质量**：
>
> 对实时性强且重要的数据报文优先处理；对于实时性不强的普通数据报文，提供较低的处理优先级，网络拥塞时甚至丢弃。QoS 应运而生。支持 QoS 功能的设备，能够提供传输品质服务；针对某种类别的数据流，可以为它赋予某个级别的传输优先级，来标识它的相对重要性，并使用设备所提供的各种优先级转发策略、拥塞避免等机制为这些数据流提供特殊的传输服务。配置了 QoS 的网络环境，增加了网络性能的可预知性，并能够有效地分配网络带宽，更加合理地利用网络资源。

[百科参考](https://baike.baidu.com/item/qos/404053)

### MQTT QoS

MQTT 设计了一套保证消息稳定传输的机制，包括消息应答、存储和重传。在这套机制下，提供了三种不同层次 QoS（Quality of Service）：

- QoS0，At most once，至多一次；-- AMO
- QoS1，At least once，至少一次；-- ALO
- QoS2，Exactly once，确保只有一次。 -- EO

QoS 是消息的发送方（Sender）和接受方（Receiver）之间达成的一个协议：

- QoS0 代表，Sender 发送的一条消息，Receiver 最多能收到一次，也就是说 Sender 尽力向 Receiver 发送消息，**如果发送失败，也就算了**；
- QoS1 代表，Sender 发送的一条消息，Receiver 至少能收到一次，也就是说 Sender 向 Receiver 发送消息，**如果发送失败，会继续重试，直到 Receiver 收到消息为止，但是因为重传的原因，Receiver 有可能会收到重复的消息**；-- 处理消息的方法做到幂等，就算消息重复也不怕了
- QoS2 代表，Sender 发送的一条消息，Receiver 确保能收到而且只收到一次，也就是说 Sender 尽力向 Receiver 发送消息，**如果发送失败，会继续重试，直到 Receiver 收到消息为止，同时保证 Receiver 不会因为消息重传而收到重复的消息**。

[参考](https://zhuanlan.zhihu.com/p/80203905)

## Go 库

[Go 实现的支持集群的高性能的 MQTT 库](https://github.com/fhmq/hmq/)

- 支持 MQTT 版本：3.1.1

### go.mod

看下它依赖了什么：

```go
module github.com/fhmq/hmq

go 1.12

require (
 github.com/Shopify/sarama v1.23.0 // Kafka的Go客户端
 github.com/bitly/go-simplejson v0.5.0
 github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
 github.com/eapache/queue v1.1.0
 github.com/eclipse/paho.mqtt.golang v1.2.0 // 另一个mqtt的golang库
 github.com/gin-gonic/gin v1.7.0
 github.com/google/uuid v1.1.1
 github.com/kr/pretty v0.1.0 // indirect
 github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
 github.com/modern-go/reflect2 v1.0.1 // indirect
 github.com/patrickmn/go-cache v2.1.0+incompatible
 github.com/pkg/errors v0.8.1 // indirect
 github.com/segmentio/fasthash v0.0.0-20180216231524-a72b379d632e
 github.com/stretchr/testify v1.4.0
 github.com/tidwall/gjson v1.9.3
 go.uber.org/atomic v1.4.0 // indirect
 go.uber.org/multierr v1.1.0 // indirect
 go.uber.org/zap v1.10.0
 golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
 gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
 gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect Standard interface for holding authenticated identities and their attributes.
)
```

### 示例

TODO:

### 源码

TODO:
