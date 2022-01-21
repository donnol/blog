package generic

// 编译时，类型检查`递归调用`时栈溢出
// 
// 使用`go version devel go1.18-d15481b Fri Jan 21 01:14:28 2022 +0000 linux/amd64`编译时，不再报溢出错误；
// 但是，会报：type int of int(0) does not match P
// 
// > With the above CL this doesn't crash anymore. Still needs a proper fix but not a release blocker anymore. [From](https://github.com/golang/go/issues/48656)
func f[P any](a, _ P) {
	f(a, int(0))
}
