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