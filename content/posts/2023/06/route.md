---
author: "jdlau"
date: 2023-06-13
linktitle: windows route
menu:
next:
prev:
title: windows route
weight: 10
categories: ['route']
tags: ['route']
---

## `Windows`通过`route`命令配置路由

```sh
route add 192.168.0.0 mask 255.255.0.0 192.168.66.254 -p
```

> 192.168.0.0: 目标主机的网络地址
>
> mask 255.255.0.0: 掩码，与目标网络地址对应
>
> 192.168.66.254: 网关地址

## `Linux`通过`ip route`命令配置路由

**NOTE**: `ip route`是`route`命令的升级版本，但route命令仍在大量使用。

```sh
# 设置192.168.4.0网段的网关为192.168.166.1,数据走wlan0接口
# /24 is the network prefix. The network prefix is the number of enabled bits in the subnet mask.
# 24位子网掩码
ip route add 192.168.4.0/24 via 192.168.166.1 dev wlan0

# 255.255.255.0为子网掩码
# 3*8(255即是8位二进制)
ip route add 192.168.0.0/255.255.255.0 dev eth0
```

## 子网掩码

`ip`地址包含了网络地址和主机地址两部分，怎么区分呢？

这是就需要用到子网掩码了，它是一个与`ip`地址同位数、连续的数，可以用位数`24`表示，也可以用地址`255.255.255.0`表示。

两个地址经过位与运算后，`ip`中没被掩盖的部分即是网络地址，被掩盖的部分即是主机地址。
