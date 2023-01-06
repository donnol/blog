---
author: "jdlau"
date: 2023-01-06
linktitle: Go Generic Join
menu:
next:
prev:
title: Go Generic Join
weight: 10
categories: ['Go', 'Generic', 'Join']
tags: ['Join']
---

```go
// nested loop, bad but common
func Join[J, K, R any](
	left []J,
	right []K,
	match func(J, K) bool,
	mapper func(J, K) R,
) []R {
	var r = make([]R, 0, len(left))

	for _, j := range left {
		for _, k := range right {
			if !match(j, k) {
				continue
			}
			r = append(r, mapper(j, k))
		}
	}

	return r
}

// hash join, good but specify -- must use equal condition
func JoinByKey[K comparable, LE, RE, R any](
	left []LE,
	right []RE,
	lk func(item LE) K,
	rk func(item RE) K,
	mapper func(LE, RE) R,
) []R {
	var r = make([]R, 0, len(left))

	rm := KeyBy(right, rk)

	for _, j := range left {
		k := lk(j)
		re := rm[k]
		r = append(r, mapper(j, re))
	}

	return r
}
```

[Code From](https://github.com/donnol/lo)
