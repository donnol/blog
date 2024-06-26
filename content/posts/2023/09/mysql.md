---
author: "jdlau"
date: 2023-09-15
linktitle: 查找并杀掉运行中事务
menu:
next:
prev:
title: 查找并杀掉运行中事务
weight: 10
categories: ['Mysql']
tags: ['note']
---

## 查找并杀掉运行中事务

```sql
-- 获取线程id然后杀掉
SELECT * FROM information_schema.innodb_trx;
-- 优先找到其中耗时最长的删
kill 36272;
kill 36275;
kill 35971;
kill 35972;

-- 其它
select * from performance_schema.events_statements_current;
show processlist;
```

## 查看锁使用情况

```sql
SELECT object_name, index_name, lock_type, lock_mode, lock_data FROM performance_schema.data_locks;
```

## Windows环境下死锁时重启数据库导致数据库出现“启动后停止”

后面重启服务器也不行，一直提示：

```sh
> net start mysql
MYSQL服务正在启动
MYSQL服务无法启动

服务没有报告任何错误。
请键入 MET HELPMSG 3534 以获得更多的帮助。
```

解决办法：

```sh
将data目录移走，再重新初始化数据库，此时可正常启动数据库；然后停止数据库，将旧data数据库里的数据库和索引文件复制回新的data目录里，再启动数据库。

# 重新初始化数据库
mysqld --initialize-insecure --user=mysql
mysqld -install

# 停止和启动
net stop mysql
net start mysql
```

[参考](https://blog.csdn.net/weixin_46483006/article/details/136692632)

其中需要复制的文件有这些：

```sh
[数据库目录]
#ib_16384_0.dblwr
#ib_16384_1.dblwr
ib_fubber_pool
ibdata1
mysql.ibd
```
