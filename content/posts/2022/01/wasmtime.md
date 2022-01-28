---
author: "jdlau"
date: 2022-01-28
linktitle: wasm运行时wasmtime
menu:
next:
prev:
title: wasm运行时wasmtime
weight: 10
categories: ['wasi', 'wasm', 'runtime']
tags: ['wasmtime']
---

## 源码

```sh
# 下载
git clone git@github.com:bytecodealliance/wasmtime.git

# 子模块
git submodule update --init --recursive

# 安装
cargo build
```

如果忘了拉子模块，`vscode`的`rust-analyzer`会报错，导致智能提示等功能失效。

不过整个初始化过程还是有点长，等了好久才能正常使用。

### 阅读

从`build.rs`开始，首先映入眼帘的是`use anyhow::Context;`：

```rust
/// Provides the `context` method for `Result`.
///
/// This trait is sealed and cannot be implemented for types outside of
/// `anyhow`.
```

这是一个为其它类型（`anyhow::Result`）引入`context`方法的特征啊，多么伟大，在`anyhow`包外面的类型就不要想着去实现它了，你们高攀不起的。

再看`anyhow::Context`的定义：

```rust
// lib.rs:598
pub trait Context<T, E>: context::private::Sealed { // 继承了Sealed，那它又是怎么样的、做什么的呢？
    /// Wrap the error value with additional context. -- 给error值包装上下文信息
    fn context<C>(self, context: C) -> Result<T, Error>
    where
        C: Display + Send + Sync + 'static; // 能展示，并发安全，全局可见的类型值

    /// Wrap the error value with additional context that is evaluated lazily
    /// only once an error does occur. -- 通过传入一个FnOnce的函数来延迟获取上下文信息
    fn with_context<C, F>(self, f: F) -> Result<T, Error>
    where
        C: Display + Send + Sync + 'static,
        F: FnOnce() -> C;
}

// context.rs:170
pub(crate) mod private { // pub(crate)表明这个mod只能在本crate里被使用
    use super::*; // 使用父mod的东西

    pub trait Sealed {} // 特征里没有方法

    impl<T, E> Sealed for Result<T, E> where E: ext::StdError {} // 为Result实现Sealed
    impl<T> Sealed for Option<T> {} // 为Option实现Sealed
}
```

### 主要逻辑

做了哪些事情呢？
