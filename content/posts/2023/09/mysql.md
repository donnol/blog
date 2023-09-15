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
