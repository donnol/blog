---
author: "jdlau"
date: 2021-12-20
linktitle: vscode-go在go.mod在非根目录情况下失效的问题
menu:
next:
prev:
title: vscode-go在go.mod在非根目录情况下失效的问题
weight: 10
---

问题如图：

![问题](/image/vscode-go-module-not-support-noroot-gomod.png)

解决：

添加配置：

```json
{
    // ...
   "gopls": {
        "experimentalWorkspaceModule": true
    },
    // ...
}
```

等`go 1.18`的`workspace`模式推出之后，应该就不需要配置这个了。

[参考](https://stackoverflow.com/questions/59732657/how-do-i-properly-use-go-modules-in-vscode)
