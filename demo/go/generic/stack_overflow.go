package generic

// 编译时，类型检查`递归调用`时栈溢出
// func f[P any](a, _ P) {
// 	f(a, int(0))
// }
