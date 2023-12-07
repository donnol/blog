---
author: "jdlau"
date: 2023-12-07
linktitle: dbeaver手动导入驱动
menu:
next:
prev:
title: dbeaver手动导入驱动
weight: 10
categories: ['DB']
tags: ['mark']
---

因为直接下载`dbeaver`的时候，是没有带上驱动文件的，所以需要在使用时下载。

但是，如果刚好安装的环境是无法通网的，那么就需要手动传入驱动并安装。

做法如下：

1. 现在本地有网环境下载驱动文件

> 用`dbeaver`下载`mysql`的驱动，会存放在目录：`C:\Users\{用户名}\AppData\Roaming\DBeaverData\drivers\maven\maven-central\mysql`.
>

NOTE: 注意替换`{用户名}`为你本机实际名称。

2. 把下好的文件传入到无网机器上，同样放到以上目录。

3. 打开`dbeaver`，`数据库`->`驱动管理器`，添加驱动

> 选中`MySQL`，然后点击编辑；在弹出框里切到`库`，将已有内容全部删掉，再点击`添加文件夹`，然后选择上面驱动存放的目录，即可确定保存。

如此，即可手动导入驱动文件。
