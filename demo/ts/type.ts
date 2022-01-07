type I = string & number; // never
type U = string | number;
type A = string;

type M = {
    a: string;
    b: number;
} | U;

function printm(m: M): U {
    console.log(m);
    return 0;
}
