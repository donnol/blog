---
author: "jdlau"
date: 2023-05-16
linktitle: mysqlrouter
menu:
next:
prev:
title: mysqlrouter使用
weight: 10
categories: ['mysql', 'router']
tags: ['mysqlrouter']
---

## What

`mysqlrouter`是一个代理，可以将查询转发到配置好的数据库服务里。

## Why

在办公室网络环境下基于`win10 wsl2`开发应用时，需要连接到主机所在局域网的其它机器上的数据库服务。

> 也就是说，存在机器：wsl2、主机、其它机器。

`wsl2`通过`NAT`网络模式与`主机`互通，并且`wsl2`可以访问外网。

但是`wsl2`不能访问到`其它机器`上的数据库服务，不知道是不是办公室网络环境存在限制。

为了使得`wsl2`能访问到`其它机器`上的数据库服务成立，在`主机`启动`mysqlrouter`充当代理，然后`wsl2`通过访问代理来访问`其它机器`。

## Install

可以使用`mysql installer`选择安装。

## 简单模式

配置文件（mysqlrouter.conf）：

```config
[DEFAULT]
logging_folder = D:/Data/mysqlrouter/log
plugin_folder = C:/Program Files/MySQL/MySQL Router 8.0/lib # 这里是插件所在目录，必须是mysqlrouter安装路径下的目录，否则报错找不到插件
config_folder = D:/Data/mysqlrouter/etc # 启动配置默认查找目录，会在目录里寻找mysqlrouter.conf文件
runtime_folder = D:/Data/mysqlrouter/run
data_folder = D:/Data/mysqlrouter/data

[logger]
level = DEBUG

[routing:primary]
bind_address=172.20.96.1 # 主机ip地址
bind_port=6446 # 主机监听端口
destinations = 172.17.39.239:3306 # 目标机器，也就是实际执行查询的数据库服务所在机器的地址
mode = read-write
connect_timeout = 10
```

启动：`mysqlrouter -c D:\Data\mysqlrouter\etc\mysqlrouter.conf`

关闭防火墙或者配置规则允许端口通过。

在`wsl2`机器上访问：`mysql -h 172.20.96.1 -P 6446 -uroot -p`，即可访问到`172.17.39.239:3306`机器上的数据库服务。

## 问题

报错：`Error: Loading plugin... 找不到指定的程序。`

```sh
mysqlrouter -c D:\Data\mysqlrouter\etc\mysqlrouter.conf
Error: Loading plugin for config-section '[routing:primary]' failed: D:/Data/mysqlrouter/lib/routing.dll: 找不到指定的程序。
PS D:\Project> 
PS D:\Project> mysqlrouter_plugin_info D:/Data/mysqlrouter/lib/routing.dll routing
[ERROR] Could not load plugin file 'D:/Data/mysqlrouter/lib/routing.dll': 找不到指定的程序。
```

解决：需要将`plugin_folder`配置的值改为`mysqlrouter安装路径下的目录`。

## 其它代理

如果只是针对`mysql`，则使用`mysqlrouter`即可，但如果还有其它服务，则还不够。

### goproxy

这时，可以使用[`goproxy`](https://github.com/snail007/goproxy)。

注意，这个项目的源码不是最新的（直接使用基于源码构建出来的`proxy`会有与文档不一致的表现），需要去[下载](https://github.com/snail007/goproxy/releases)最新的二进制执行文件.

解压后执行，即可启动代理，如下述命令将启动一个代理，其监听本地`33080`端口，并将请求转发到`172.17.39.239:3306`目标机器：

```sh
.\proxy.exe tcp -p ":33080" -T tcp -P "172.17.39.239:3306"
```

### 访问sqlserver

通过代理(不同的端口)连接`172.17.39.239`机器上的`sqlserver`:

```sh
.\proxy.exe tcp -p ":33081" -T tcp -P "172.17.39.239:1433"
```

### 加密

如果要生成密钥，必须在`linux`环境下使用`proxy keygen`生成密钥(在`windows`环境里会报错)。

### 自定义tcp代理

[地址](`github.com/donnol/do/cmd/letgo`).

> 下载源码：git clone git@github.com:donnol/do.git
>
> 安装: go install ./cmd/letgo
>
> 使用: letgo proxy --remoteAddr="172.17.39.239:3306"

默认监听端口`54388`，将收到的数据转发到指定的地址：`172.17.39.239:3306`。

可通过监听不同的端口转发数据到不同的地址。
