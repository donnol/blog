---
author: "jdlau"
date: 2022-07-20
linktitle: Jump Table
menu:
next:
prev:
title: Jump Table
weight: 10
categories: ['Go']
tags: ['version']
---

背景：分多次把一批货全部出清。

要求：需要确保这批货多次出清跟一次出清收的钱一样。

现有三个数字(可整数，可小数)：a b c，其中：a为数量，b为价格，c为折扣。

总额为: t, `t = a*b*c`

引入中间量: x y z

```
x = a
y = x*b
z = y*c
```

假设分三次，每次数量为：a1 a2 a3，则有：a = a1 + a2 + a3

第1次.

```
x1 = (x-a1)
y1 = (y-y*a1/x)
z1 = (z-z*a1/x)
t1 = z*a1/x
```

第2次.

```
x2 = (x1-a2)
y2 = (y1-y1*a2/x1)
z2 = (z1-z1*a2/x1)
t2 = z1*a2/x1
```

第3次.

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
