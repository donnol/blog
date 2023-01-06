type I = string & number; // never
type U = string | number;
type A = string;

type M = {
    a: string;
    b: number;
} | U;

function printm(m: M): U {
    console.log(m);
    let type = typeof m;
    console.log(type);
    return 0;
}

printm(212);
printm("abc");

type MK<T> = T extends void ? {
    more?: boolean
} : {
    more?: boolean
    params: T
}

export class MultiPager<T1, T2 = void> {
    fetch = (opts: MK<T2>) => {
        console.log(opts)
    }
}

new MultiPager<string>().fetch({ more: true })
new MultiPager<string, number>().fetch({ more: true, params: 1 })

export class MultiPager2<T1, T2 = void> {
    fetch = (opts: { more?: boolean, params?: T2 }) => { // 用params?会导致范围变大
        console.log(opts)
    }
}

new MultiPager2<string>().fetch({ more: true })
new MultiPager2<string>().fetch({ more: true, params: void (1) })
new MultiPager2<string, number>().fetch({ more: true })
new MultiPager2<string, number>().fetch({ more: true, params: 1 })

// 根据传入的类型，选择返回的类型
// 根据传入的类型，提取它里面的子类型并返回

// without infer
type Ids = number[];
type Names = string[];

type Unpacked<T> = T extends Names ? string : T extends Ids ? number : T;

type idType = Unpacked<Ids>; // idType 类型为 number
type nameType = Unpacked<Names>; // nameType 类型为string

// with infer
type Unpacked2<T> = T extends (infer R)[] ? R : T;

type idType2 = Unpacked<Ids>; // idType 类型为 number
type nameType2 = Unpacked<Names>; // nameType 类型为string
