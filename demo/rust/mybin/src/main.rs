pub mod tree;
use crate::tree::tree::{new_tree, Treer};

#[tokio::main]
async fn main() {
    // 要想在main使用async，必须在上面加上tokio::main属性，并且引入带"rt-multi-thread", "macros"特性的tokio库
    println!("Hello, world!");

    println!("{}", add(1, 2));
    println!("{}", add1(1, 2));

    // const fn set value to a const var
    const V: u8 = add(2, 4);
    println!("{}", V);

    // error[E0015]: calls in constants are limited to constant functions, tuple structs and tuple variants
    // const v1: u8 = add1(4, 3);
    // println!("{}", v1)

    println!("{}", pi(10));

    let r = y().await;
    println!("{}", r);

    // generator_yield();

    let t = new_tree();
    println!("{:?}", t);
    let value = t.find("1".to_string());
    println!("{:?}", value);

    let mut tree1 = new_tree();
    tree1.insert("1".to_string(), "4".to_string());
    let value = tree1.find("1".to_string());
    println!("{:?}", value);

    println!("block: {}", block());
    println!("join: {:?}", block_on(join()));
}

const fn add(x: u8, y: u8) -> u8 {
    x + y
}

fn add1(x: u8, y: u8) -> u8 {
    x + y
}

// https://en.wikipedia.org/wiki/Bailey–Borwein–Plouffe_formula
fn bbp(k: u32) -> f64 {
    let a1 = 4.0 / (8 * k + 1) as f64;
    let a2 = 2.0 / (8 * k + 4) as f64;
    let a3 = 1.0 / (8 * k + 5) as f64;
    let a4 = 1.0 / (8 * k + 6) as f64;

    (a1 - a2 - a3 - a4) / ((16 as f64).powi(k as i32))
}

pub fn pi(n: u32) -> f64 {
    let mut result: f64 = 0.0;
    for i in 0..n {
        result += bbp(i);
    }
    result
}

async fn x() -> usize {
    5
}

// 将上面的x替换为下面的async_x
use std::future::Future;
use futures::executor::block_on;

#[allow(dead_code)]
#[inline(never)]
fn async_x() -> impl Future<Output = usize> {
    // 使用了async后，返回的其实是一个实现了Future trait的对象
    async { 5 }
}

async fn y() -> usize {
    let r = x().await; // 出现await时，当函数执行被阻塞时，会将执行权交出，并且会有轮询在后台运行，直到值返回，才继续执行后续逻辑

    println!("{}", "y complete");

    r
}

fn block() -> usize {
    let r = x();
    block_on(r)
}

async fn join() -> (usize, usize) {
    futures::join!(x(), x())
}

// From https://mp.weixin.qq.com/s/ZGuqqFOcoUERMnGMtpNuIA
// 报错：error[E0658]: yield syntax is experimental

// https://doc.rust-lang.org/beta/unstable-book/language-features/generators.html
// 加了下面这句也不行
// #![feature(generators, generator_trait)]

// use std::ops::{Generator, GeneratorState};
// use std::pin::Pin;

// fn generator_yield() {
//     let mut generator = || {
//         let mut val = 1;
//         yield val; // 通过yield来保存上下文，并离开现场，下次恢复执行时，再次执行接下来的逻辑
//         val += 1;
//         yield val;
//         val += 1;
//         val
//     };

//     match Pin::new(&mut generator).resume(()) {
//         GeneratorState::Yielded(1) => {}
//         _ => panic!("unexpected value from resume"),
//     }
//     match Pin::new(&mut generator).resume(()) {
//         GeneratorState::Yielded(2) => {}
//         _ => panic!("unexpected value from resume"),
//     }
//     match Pin::new(&mut generator).resume(()) {
//         GeneratorState::Complete(3) => {}
//         _ => panic!("unexpected value from resume"),
//     }
// }

#[allow(dead_code)]
fn fnd() -> impl Fn() -> usize {
    || 1

    // core::ops::function
    // pub trait Fn<Args>
    // where
    //     Self: FnMut<Args>,
    //
    // #[fundamental] // so that regex can rely that `&str: !FnMut`
    // #[must_use = "closures are lazy and do nothing unless called"]
    // pub trait Fn<Args>: FnMut<Args> { - 继承了FnMut，而FnMut又继承了FnOnce；
    //     /// Performs the call operation.
    //     #[unstable(feature = "fn_traits", issue = "29625")]
    //     extern "rust-call" fn call(&self, args: Args) -> Self::Output;
    // }
    //
    // The version of the call operator that takes an immutable receiver.
    // - 不可变receiver
    // Instances of Fn can be called repeatedly without mutating state.
    // - Fn实例可以在不改变状态情况下被重复调用
    // This trait (Fn) is not to be confused with [function pointers] (fn).
    // - 不要与函数指针混淆
    // Fn is implemented automatically by closures which only take immutable references to captured variables or don't capture anything at all, as well as (safe) [function pointers] (with some caveats, see their documentation for more details). Additionally, for any type F that implements Fn, &F implements Fn, too.
    // - 捕获的是不可变引用的变量或没有捕获任何变量的闭包自动实现Fn。另外，任何实现了Fn的类型F，&F也实现了Fn。

    // Since both FnMut and FnOnce are supertraits of Fn, any instance of Fn can be used as a parameter where a FnMut or FnOnce is expected.
    // - 因为FnMut和FnOnce都是Fn的超特征，任何Fn的实例都能被用作FnMut或FnOnce。
}

#[allow(dead_code)]
fn fnmutd() -> impl FnMut() -> usize {
    || 1

    // core::ops::function
    // pub trait FnMut<Args>
    // where
    //     Self: FnOnce<Args>,
    //
    // #[fundamental] // so that regex can rely that `&str: !FnMut`
    // #[must_use = "closures are lazy and do nothing unless called"]
    // pub trait FnMut<Args>: FnOnce<Args> {
    //     /// Performs the call operation.
    //     #[unstable(feature = "fn_traits", issue = "29625")]
    //     extern "rust-call" fn call_mut(&mut self, args: Args) -> Self::Output;
    // }
    //
    // The version of the call operator that takes a mutable receiver.
    // - 可变receiver
    // Instances of FnMut can be called repeatedly and may mutate state.
    // - 可以被重复调用，可能修改状态
    // FnMut is implemented automatically by closures which take mutable references to captured variables, as well as all types that implement Fn, e.g., (safe) [function pointers] (since FnMut is a supertrait of Fn). Additionally, for any type F that implements FnMut, &mut F implements FnMut, too.
    // - 捕获了可变引用的变量，或者实现了Fn的所有类型，都会自动实现FnMut。
    // Since FnOnce is a supertrait of FnMut, any instance of FnMut can be used where a FnOnce is expected, and since Fn is a subtrait of FnMut, any instance of Fn can be used where FnMut is expected.
    // 因为FnOnce是FnMut的超特征，任何实现了FnMut的实例都可以被用作FnOnce。
}

#[allow(dead_code)]
fn fnonced() -> impl FnOnce() -> usize {
    || 1

    // core::ops::function
    // pub trait FnOnce<Args>
    //
    // #[rustc_on_unimplemented(
    //     on(
    //         Args = "()", - Args是一个小括号？
    //         note = "wrap the `{Self}` in a closure with no arguments: `|| {{ /* code */ }}`"
    //     ),
    //     message = "expected a `{FnOnce}<{Args}>` closure, found `{Self}`",
    //     label = "expected an `FnOnce<{Args}>` closure, found `{Self}`"
    // )]
    // #[fundamental] // so that regex can rely that `&str: !FnMut`
    // #[must_use = "closures are lazy and do nothing unless called"]
    // pub trait FnOnce<Args> { - Args没有指定约束，但是在上面的rustc_on_unimplemented出现了：`Args = "()"`
    //     /// The returned type after the call operator is used.
    //     #[lang = "fn_once_output"]
    //     #[stable(feature = "fn_once_output", since = "1.12.0")]
    //     type Output; - 关联类型，作为返回值类型

    //     /// Performs the call operation.
    //     #[unstable(feature = "fn_traits", issue = "29625")]
    //     extern "rust-call" fn call_once(self, args: Args) -> Self::Output;
    // }
    //
    // The version of the call operator that takes a by-value receiver.

    // Instances of FnOnce can be called, but might not be callable multiple times. Because of this, if the only thing known about a type is that it implements FnOnce, it can only be called once.
    // - FnOnce的实例不能被调用多次。只能被调用一次
    // FnOnce is implemented automatically by closures that might consume captured variables, as well as all types that implement FnMut, e.g., (safe) [function pointers] (since FnOnce is a supertrait of FnMut).
    // - 使用捕获变量的闭包自动实现FnOnce。
    // Since both Fn and FnMut are subtraits of FnOnce, any instance of Fn or FnMut can be used where a FnOnce is expected.
    // -
    // Use FnOnce as a bound when you want to accept a parameter of function-like type and only need to call it once. If you need to call the parameter repeatedly, use FnMut as a bound; if you also need it to not mutate state, use Fn.
    // - 当需要接收函数类型参数并只调用它一次时使用FnOnce。如果想调用多次，使用FnMut；如果还要它不改变状态，使用Fn。
}
