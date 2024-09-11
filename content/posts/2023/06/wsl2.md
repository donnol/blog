---
author: "jdlau"
date: 2023-06-01
linktitle: wsl2初始化Mysql数据库速度非常慢
menu:
next:
prev:
title: wsl2初始化Mysql数据库速度非常慢
weight: 10
categories: ['mysql', 'wsl2']
tags: ['mysql']
---

## wsl2版本

```sh
> wsl.exe -v
WSL version: 1.2.5.0
Kernel version: 5.15.90.1
WSLg version: 1.0.51
MSRDC version: 1.2.3770
Direct3D version: 1.608.2-61064218
DXCore version: 10.0.25131.1002-220531-1700.rs-onecore-base2-hyp
Windows version: 10.0.19045.3031
```

使用过程中，因为磁盘空间问题，把子系统安装位置从`C盘`转移到了其它盘。

## 操作

把`sql`目录里的`*.sql`文件逐一导入到`8.0.33`版本的`Mysql`。

尽管`sql`文件不多也不大，但是整个过程非常慢。其中一个有一千个左右的`INSERT IGNORE`语句，更是用了将近12分钟才完成。

## 怎么办？

改为通过网络访问本机的数据库。

## 端口

[一文中](https://zhuanlan.zhihu.com/p/425312804)提到：

> 根据我的观察, 如果Windows本地启动了指定端口, 这时WSL2中虽然可以使用相同的端口, 但是localhost:port 将指向Windows的服务, WSL的服务将会被覆盖!
>
> 当然了, 如果我们配置了端口转发, 转发的IP是WSL的地址, 而不是localhost, 那么WSL将会覆盖Windows的服务!

而我的观察是，

1. 我发现本地起了数据库服务之后，在wsl2里起数据库服务（mysql服务，端口都是3306）的情况下，是不会报端口重复绑定错误的。

2. 但是如果我在wsl2里先起一个服务，绑定端口14222后，再在主机起相同服务，想绑定相同端口时，则会报错端口已被绑定。

3. 如果我是先在主机起上述服务，然后再在wsl2起该服务，则能正常启动。那在主机访问`localhost:[port]`时会访问到哪个呢？此时访问到的是主机的服务。

所以端口是否占用会不会还跟**服务起的顺序**有关呢？

暂时未看到有确切的描述。

但是，经过上面的实验，可以认为：

    先在主机起服务，再在wsl2起服务绑定相同端口时，服务可正常启动；在主机访问`localhost:[port]`时访问的是主机的服务；在wsl2里访问的则是wsl2的服务，除非手动指定主机IP。

    而如果先在wsl2里起了服务，再在主机起服务（绑定相同端口: 14222），则会报错端口已被绑定。

    但是，如果在wsl2里先起的mysql服务，再在主机起，则不会报错，所以还跟端口值有关？

### 最新发现

如何通过局域网访问WSL2中的服务？

假设局域网上有两台主机A和B。主机A安装了WSL2、开启了Redis服务，端口为6379。现在主机B如何才能访问主机A上WSL2的Redis服务呢?

1. 配置端口转发

1). 以管理员权限打开PS，输入命令：

```sh
# listenaddress： 监听地址， 0.0.0.0 表示匹配所有地址。
# listenport：监听的Windows端口。
# connectaddress：要转发的地址。-- 下面的`172.20.109.210`是`WSL2`的ip地址
# connectport： 转发的WSL2端口。
netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=6379 connectaddress=172.20.109.210 connectport=6379
```

2). 通过以下命令，查看当前所有的转发设置。

```sh
netsh interface portproxy show all
```

也可以通过以下命令来删除转发设置：

```sh
netsh interface portproxy delete v4tov4 listenaddress=0.0.0.0 listenport=6379
```

2. 防火墙配置入站规则

防火墙->防火墙和网络保护->高级设置->入站规则->新建规则->填写允许通过的端口完成添加

出站规则同上。

### 最后

WSL2启动服务后，本机可以直接使用`localhost:6379`访问到该服务，但是用`127.0.0.1:6379`不行。

而做完以上`netsh`转发操作后，本机访问`127.0.0.1:6379`时报错：`ERR_EMPTY_RESPONSE`，与期望的结果不符，这是为什么呢？

[WSL2网络配置的讨论](https://github.com/microsoft/WSL/issues/4150)

最后的最后，重启了电脑之后，突然就可以了。

#### WSL2 2.0新增实验配置：`networkingMode=mirrored`，但是需要Win11才能用。

[WSL2 和 Windows 主机的网络互通而且 IP 地址相同了，还支持 IPv6 了，并且从外部（比如局域网）可以同时访问 WSL2 和 Windows 的网络。](https://www.v2ex.com/t/975098?p=2)
