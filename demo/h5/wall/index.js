// export default 当文件被使用script直接引入到html文件时，不能使用，否则报错：Uncaught SyntaxError: export declarations may only appear at top level of a module
// 可以在script里设置type="module"来消除这个错误，但又带来另外的问题，onclick的时候会找不到printName方法

function printName() {
    console.log("wall")
}
