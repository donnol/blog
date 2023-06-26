---
author: "jdlau"
date: 2023-06-26
linktitle: ml sklearn
menu:
next:
prev:
title: ml sklearn
weight: 10
categories: ['sklearn']
tags: ['ml']
---

## 尝试

```sh
$ pip install scikit-learn
$ python3.10
>>> from sklearn.datasets import load_iris
>>> from sklearn.linear_model import LogisticRegression
>>> data, y = load_iris(return_X_y=True)
>>> clf = LogisticRegression(random_state=0, max_iter=1000).fit(data, y)
>>> clf.predict(data[:2, :])
>>> clf.predict_proba(data[:2, :])
>>> clf.score(data, y)
```

### 查找操作记录

`cat ~/.python_history`.
