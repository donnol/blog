console.log("9007199254740991".length)
// 16
// undefined
console.log(Number.MAX_SAFE_INTEGER)
// 9007199254740991
// undefined
console.log(Number.MIN_SAFE_INTEGER)
// -9007199254740991
// undefined
console.log(Number.MAX_SAFE_INTEGER+1) // +1 和 +2 得到的结果居然一样
// 9007199254740992
// undefined
console.log(Number.MAX_SAFE_INTEGER+2)
// 9007199254740992
// undefined
console.log(Number.MAX_SAFE_INTEGER+3)
// 9007199254740994
// undefined
console.log(Math.pow(2, 53) - 1)
// 9007199254740991
