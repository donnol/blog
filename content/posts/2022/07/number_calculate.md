---
author: "jdlau"
date: 2022-07-20
linktitle: 数字计算之分摊
menu:
next:
prev:
title: 数字计算之分摊
weight: 10
categories: ["number"]
tags: ["calculate"]
---

背景：分多次把一批货全部出清。

要求：需要确保这批货多次出清跟一次出清收的钱一样。

现有三个数字(可整数，可小数)：a b c，其中：a 为数量，b 为价格，c 为折扣。

则总额为: t, `t = a*b*c`

假设分三次，每次数量为：a1 a2 a3，则有：a = a1 + a2 + a3

1. 直接计算：

第 1 次.

```
a1*b*c
```

第 2 次.

```
a2*b*c
```

第 3 次.

```
a3*b*c
```

(a1+a2+a3)*b*c 不就等于 a*b*c 了吗？

但是，如果考虑到小数乘法计算时的精度，比如：1.22*2.33 相乘后再取精度（保留两位小数），不就会导致数量误差了吗？

那如果取精度导致结果误差，那我不取精度，直接用所有小数位数来计算呢。

虽说可以，但小数位数是有可能非常多的，占用的空间也是一笔不小的开销。

2. 引入中间量(可称为'余额'): x y z

```
x = a
y = x*b
z = y*c
```

第 1 次.

```
x1 = (x-a1)
y1 = (y-y*a1/x)
z1 = (z-z*a1/x)
t1 = z*a1/x
```

第 2 次.

```
x2 = (x1-a2)
y2 = (y1-y1*a2/x1)
z2 = (z1-z1*a2/x1)
t2 = z1*a2/x1
```

第 3 次.

```
x3 = (x2-a3)
y3 = (y2-y2*a3/x2)
z3 = (z2-z2*a3/x2)
t3 = z2*a3/x2
```

求证：t = t1 + t2 + t3 ?

```
a*b*c =
    (a*b*c*a1/a) +

 ((a*b*c - a*b*c*a1/a)*a2/(a-a1)) +

    ((a*b*c - a*b*c*a1/a) - (a*b*c - a*b*c*a1/a)*a2/(a-a1))*a3/(a-a1-a2)

a*b*c =
    (a*b*c*a1/a) +

 ((a*b*c - a*b*c*a1/a)*a2/(a-a1)) +

    ((a*b*c - a*b*c*a1/a) - (a*b*c - a*b*c*a1/a)*a2/(a-a1))

a*b*c =
    (a*b*c*a1/a) +

 ((a*b*c - a*b*c*a1/a)*a2/(a-a1)) +

    -- (a*b*c - a*b*c*a1/a)*(1 - a2/(a-a1))
    (a*b*c - a*b*c*a1/a)*(a3/(a2+a3))

a*b*c =
    (a*b*c*a1/a) +

 (a*b*c - a*b*c*a1/a)*a2/(a2+a3) +

    (a*b*c - a*b*c*a1/a)*(a3/(a2+a3))

a*b*c =
    (a*b*c*a1/a) +

 (a*b*c - a*b*c*a1/a)

a*b*c = a*b*c
```
