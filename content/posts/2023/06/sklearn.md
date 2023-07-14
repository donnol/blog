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

## 模型、策略、优化算法

模型是输入输出函数：Y = F(X).

策略是拟合过程的损失函数：L(Y, F(X)), 可以是均方误差、对数损失函数、交叉熵损失函数。

优化算法：确定模型和损失函数后，可以加速计算的方法，比如：随机梯度下降法、牛顿法、拟牛顿法。
