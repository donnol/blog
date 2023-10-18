---
author: "jdlau"
date: 2023-10-18
linktitle: `Mysql 8.0.33`在使用窗口函数的同时不能用`count(distinct *)`
menu:
next:
prev:
title: `Mysql 8.0.33`在使用窗口函数的同时不能用`count(distinct *)`
weight: 10
categories: ['Mysql 8']
tags: ['mark']
---

`Mysql 8.0.33`在使用窗口函数的同时不能用`count(distinct *)`

比如，我想在窗口函数里使用字段`apply_unit_id`分组，然后求`project_id`列不重复值的数量：

```sql
select distinct apply_unit_id, count(distinct project_id) over (partition by apply_unit_id)
from weia join weiag on weiag.apply_id = weia.id 
;
```

此时报错：`SQL 错误 [1235] [42000]: This version of MySQL doesn't yet support '<window function>(DISTINCT ..)'`

## 怎么办呢？

使用[`dense_rank()`](https://dev.mysql.com/doc/refman/8.0/en/window-function-descriptions.html#function_dense-rank)间接计算：

```sql
select distinct apply_unit_id, dense_rank() over (partition by apply_unit_id order by project_id) 
+ dense_rank() over (partition by apply_unit_id order by project_id desc) 
- 1
from weia join weiag on weiag.apply_id = weia.id 
;
```

## dense_rank()

> Returns the rank of the current row within its partition, without gaps.
>
> -- 返回当前行在其分区内的排名，没有间隙。

而[`rank()`](https://dev.mysql.com/doc/refman/8.0/en/window-function-descriptions.html#function_rank)

> Returns the rank of the current row within its partition, with gaps.
>
> -- 返回当前行在其分区内的排名，带有间隙。

那么，这里的间隙是什么意思呢？

```sh
mysql> SELECT
         val,
         ROW_NUMBER() OVER w AS 'row_number',
         RANK()       OVER w AS 'rank',
         DENSE_RANK() OVER w AS 'dense_rank'
       FROM numbers
       WINDOW w AS (ORDER BY val);
+------+------------+------+------------+
| val  | row_number | rank | dense_rank |
+------+------------+------+------------+
|    1 |          1 |    1 |          1 |
|    1 |          2 |    1 |          1 |
|    2 |          3 |    3 |          2 |
|    3 |          4 |    4 |          3 |
|    3 |          5 |    4 |          3 |
|    3 |          6 |    4 |          3 |
|    4 |          7 |    7 |          4 |
|    4 |          8 |    7 |          4 |
|    5 |          9 |    9 |          5 |
+------+------------+------+------------+
```

上面的例子比较了`row_number()`, `rank()`, `dense_rank()`三种函数的效果，从中可以看出：

```
row_number(): 从1开始单调递增，不会出现重复值
rank(): 会存在相同排名，值在增大时会出现空隙，比如在1存在两个时，会从1跳到3
dense_rank(): 会存在相同排名，值在增大时不会出现空隙，即使1存在两个时，后面的排名值也不会跳跃
```
