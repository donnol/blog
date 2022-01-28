---
author: "jdlau"
date: 2022-01-26
linktitle: 容器镜像加密
menu:
next:
prev:
title: 容器镜像加密
weight: 10
categories: ['container', 'docker', 'image']
tags: ['encrypt']
---

如果我在创建镜像时把源码也打包了进去，要怎么防止别人通过这个镜像把源码给窃取了呢？

- 加密

    镜像加密

    源码加密：在COPY源码进去之前先加密；这种适合服务器不是自己的，并且在局域网里的（接过医院系统的应该都懂吧）；留这样一份加密源码也只是在方便有bug时可以快速修复的同时，还可以稍微保护一下源码；

        先使用zip压缩源码：`zip -q -r code.zip ./code`；
        再使用gpg加密：`gpg --yes --batch --passphrase=123 -c ebpf.zip`； -- 通过`--yes --batch --passphrase`三个选项避免键盘交互，最后生成`ebpf.zip.gpg`。
        后续进入容器后，使用gpg解密：`gpg -o ebpf2.zip -d ebpf.zip.gpg`；
        再使用unzip解压：`unzip -d ebpf2 ebpf2.zip`。

在镜像构建后，还要防止`docker history -H cb0b42c0cb03 --no-trunc=true`查看镜像构建历史时，泄露秘钥等信息。-- 可使用多阶段构建：在前一阶段使用密钥加密源码，后一阶段复制加密源码，从而避免密钥泄露。因为一般只需要把后一阶段构建出来的镜像分发出去就好了，而查看后一阶段构建出来的镜像的构建历史，是看不到密钥信息的（查看前一阶段的构建历史才会看到）。

## dockerfile COPY before mkdir will get a `no such file or directory` error

~~error:~~

~~```dockerfile~~
~~# ...~~

~~RUN mkdir -p /abc~~

~~COPY --from=builder /opt/efg /abc/efg~~
~~```~~

~~没有指定创建`/abc/efg`目录，会导致后续想读取该目录内容时报错：`no such file or directory`~~

~~success:~~

~~```dockerfile~~
~~# ...~~

~~RUN mkdir -p /abc~~
~~RUN mkdir -p /abc/efg~~

~~COPY --from=builder /opt/efg /abc/efg~~
~~```~~

~~必须指定创建`/abc/efg`目录，并且要以可能出现的最长路径来创建。~~

file exist in container but can't read by go

`终于知道了，搞了一天的镜像，原来问题出在了docker-compose配置里挂载了那个路径，把COPY进去的文件覆盖了，所以一直找不到文件--配置的本地目录里没有东西~~`

## 走捷径，取巧

而忘了正路
