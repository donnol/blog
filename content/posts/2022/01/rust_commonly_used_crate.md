---
author: "jdlau"
date: 2022-01-10
linktitle: rust commonly used crate
menu:
next:
prev:
title: Rust常用库
weight: 10
categories: ['rust']
tags: ['crate']
---

## crossbeam

[crossbeam](https://github.com/crossbeam-rs/crossbeam): **Tools for concurrent programming in Rust**.

- Atomics

- Data structures

- Memory management: epoch

- Thread synchronization: channel, Parker, ShardedLock, WaitGroup

- Utilities

### channel example

```rust
use crossbeam_channel::unbounded;

let (s, r) = unbounded();

s.send("Hello, world!").unwrap();

assert_eq!(r.recv(), Ok("Hello, world!"));
```

`unbounded(无限) channel`发送时不用等接收端就绪。

另外还有`bounded channel`可在新建时指定容量，后续发送的消息数不能超过该数据 -- 除非中间有消息被取走了

当`bounded channel`容量设为0时，发送前必须等接收端就绪，一般可用于线程间等待。

[更多介绍](https://docs.rs/crossbeam/0.8.1/crossbeam/channel/index.html)

[与标准库的`sync::mpsc`对比](https://blog.csdn.net/u012067469/article/details/108544104)

### epoch

[Pin 做了什么？](https://rustmagazine.github.io/rust_magazine_2021/chapter_8/rust-lockfree-part2.html)


> crossbeam在实现无锁并发结构时，采用了基于代的内存回收方式1，这种算法的内存管理开销和数据对象的数量无关，只和线程的数量相关，因此在 以上模型中可以表现出更好的一致性和可预测性。不过Rust中的所有权系统已经保证了内存安全，那为什么还需要做额外的内存回收呢？这个问题的关键点 就在要实现**无锁并发**结构。如果使用标准库中的Arc自然就不会有内存回收的问题，但对Arc进行读写是需要锁的。

[crossbeam-channel文章](https://xiaopengli89.github.io/posts/crossbeam-channel/)

## digest

[This crate provides traits which describe functionality of cryptographic hash functions and Message Authentication algorithms.](https://docs.rs/digest/latest/digest/)

加密哈希函数和消息认证算法。

## sha2raw

[聚焦在固定大小块的sha256实现。](https://crates.io/crates/sha2raw)

## rand

[随机数生成](https://crates.io/crates/rand)

## rayon

[Rayon: A data parallelism library for Rust ](https://github.com/rayon-rs/rayon)

数据并行，很容易就能转换一系列计算到并行。保证无数据竞争。

```rust
use rayon::prelude::*;
fn sum_of_squares(input: &[i32]) -> i32 {
    // input.iter()
    input.par_iter() // <-- iter -> par_iter
         .map(|&i| i * i)
         .sum()
}
```

`Parallel iterators`小心决定如何分解数据到任务里；它会动态调整以获得最大性能。

## serde & serde_json

[serde](https://github.com/serde-rs/serde)

高效的、通用的序列化和反序列化Rust数据结构框架。

```rust
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    let point = Point { x: 1, y: 2 };

    // Convert the Point to a JSON string.
    let serialized = serde_json::to_string(&point).unwrap();

    // Prints serialized = {"x":1,"y":2}
    println!("serialized = {}", serialized);

    // Convert the JSON string back to a Point.
    let deserialized: Point = serde_json::from_str(&serialized).unwrap();

    // Prints deserialized = Point { x: 1, y: 2 }
    println!("deserialized = {:?}", deserialized);
}
```

为啥`serde_json`不用导入也能用呢？

## bellperson

[a crate for building zk-SNARK circuits](https://docs.rs/bellperson/latest/bellperson/)

zk-SNARK: 零知识简洁的非交互式知识论证，一种新颖的零知识密码学形式。是一种证明结构，在这种结构中，人们可以证明拥有某些信息，例如一个秘密秘钥，而**无需透漏该信息**，并且之间没有任何交互证明者和验证者。

零知识证明，允许一方（证明者）向另一方（验证者）证明一个陈述是真实的，而不会透漏超出陈述本身有效性的任何信息。例如，给定一个**随机数**的哈希值，证明者可以说服验证者确实**存在**一个具有该哈希值的**数字**，而无需透漏**它**是什么。

证明者不仅可以说服验证者该数字存在，而且他们实际上知道这样一个数字。-- 验证者不知道具体的数字值，但是知道它存在。

“简洁”的零知识证明可以在几毫秒内得到验证，即使是关于非常大的程序的语句，证明长度也只有几百字节。

- 在第一个零知识协议中，证明者和验证者必须来回通信多轮，

- 但在“非交互式”结构中，**证明由从证明者发送到验证者的单个消息组成**。目前，生成非交互式且足够短以发布到区块链的零知识证明的最有效的已知方法是**具有初始设置阶段**，该阶段**生成在证明者和验证者之间共享的公共参考字符串**。我们将这个公共引用字符串称为系统的**公共参数**。

> Zcash是zk-SNARKs的第一个广泛应用，Zcash强大的隐私保证源于这样一个事实，即Zcash中的屏蔽交易可以在区块链上完全加密，但仍然可以通过zk-SNARK证明在网络的共识规则下验证其有效性。

[参照](https://www.bcskill.com/index.php/archives/1192.html)

## log

[日志门面库](https://docs.rs/log/latest/log/)

> 所谓门面，其实就是它定义了一套统一的日志trait API， 抽象出来日志的常规操作，具体的日志库实现它定义的API。

[介绍文章](https://colobu.com/2019/09/22/rust-lib-per-week-log/)

Log trait是核心，它定义了三个方法：

- `fn enabled(&self, metadata: &Metadata) -> bool`: 返回这条log是否允许输出日志, 具体的日志库可以根据Metadata中的日志级别来判断

- `fn log(&self, record: &Record)`: 记录这条日志，这里日志使用Record来表示这条日志

- `fn flush(&self)`: flush缓存的日志

## anyhow

[一个基于trait对象错误类型的更容易惯用的错误处理库](https://docs.rs/anyhow/latest/anyhow/)

use `Result<T, anyhow::Error>` or `anyhow::Result<T>`

> Within the function, use ? to easily propagate any error that implements the std::error::Error trait.
>
> -- 在函数里，使用`?`传播任何实现了`Error trait`的错误

```rust
use anyhow::Result;

fn get_cluster_info() -> Result<ClusterMap> {
    let config = std::fs::read_to_string("cluster.json")?;
    let map: ClusterMap = serde_json::from_str(&config)?;
    Ok(map)
}
```

## thiserror

[thiserror](https://docs.rs/thiserror/latest/thiserror/)是方便大家为自定义的错误使用宏实现std::error::Error而设计的。

[thiserror & anyhow 文章](https://rustcc.cn/article?id=6dcbf032-0483-4980-8bfe-c64a7dfb33c7)

## num_cpus

[确定当前系统上可用的CPU数](https://docs.rs/num_cpus/latest/num_cpus/)

```rust
let cpus = num_cpus::get();
```

可以根据获取到的`CPU`数值来设置`rayon::Threadpool`。

## hex

[hex: 编码和解码16进制字符串](https://docs.rs/hex/latest/hex/)

```rust
let hex_string = hex::encode("Hello world!");

println!("{}", hex_string); // Prints "48656c6c6f20776f726c6421"
```

内部使用了`serde`库。

## bincode

[使用一个小二进制序列来编码和解码](https://docs.rs/bincode/latest/bincode/)

把一个对象转为字节序列

## byteorder

[字节序：大端或小端](https://docs.rs/byteorder/latest/byteorder/)

提供了方便的方法，在大端或小端情况下编解码数字。

## lazy_static

[延迟初始化static常量](https://docs.rs/lazy_static/latest/lazy_static/)

> Rust 静态项是一种“全局变量”。它们类似于常量，但静态项不内联使用。这意味着每个值只对应一个实例， 并且在内存中只有一个固定的地址。
>
> 静态类型活在程序的整个生命周期，只有在程序退出的时候静态项才会调用drop。
>
> 静态类型是可变的， 你可以使用 mut 关键字声明可变性。
>
> 此外，任何存储在 static 的类型都必须是 Sync。
>
> 常量和静态常量都要求给他们一个值。并且他们可能只被赋予一个值，这个值是一个常数表达式。
>
> 很多情况下，我们希望**延迟初始化静态量，只有在第一次访问的时候，或者在某个特定的时候才初始化它**，那么就可以使用lazy_static。

lazy_static提供了一个宏lazy_static!，使用这个宏把你的静态变量“包裹”起来就可以实现**延迟初始化**了。

> 实际上这个宏会帮助你生成一个特定的struct,这个struct的deref方法(trait Deref)提供了延迟初始化的能力，它也提供了initialize方法，你也可以在代码中主动地调用它进行初始化。

[更多](https://colobu.com/2019/09/08/rust-lib-per-week-lazy-static/)

## libc

[绑定到平台的系统库的FFI](https://docs.rs/libc/latest/libc/)

## pairing

[成对友好的曲线](https://docs.rs/pairing/latest/pairing/)

## blstrs

[BLS12-381成对椭圆曲线算法实现](https://docs.rs/blstrs/latest/blstrs/)
