fn main() {
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
