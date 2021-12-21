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
}

const fn add(x: u8, y: u8) -> u8 {
    x+y
}

fn add1(x: u8, y: u8) -> u8 {
    x+y
}
