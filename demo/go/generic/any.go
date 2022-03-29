package generic

// 当类型都是any时，输入参数的类型与输出参数的类型没有任何关系；
// 不管传进来的是int, string，还是其它，返回值都可以是任何类型
// 在运行时知道具体类型
func UseAnyDirect(a, b any) any {

	return true
}

// 当使用any作为类型参数，约束参数的类型时，输入参数的类型和输出参数的类型都必须遵循约束；
// 如果传进来的是int, string，返回的只能是string类型，也就是参数b的类型和结果类型必须一致；
// 从而确保接口的类型安全
// 在编译时知道具体类型
func UseAnyGeneric[T, R any](a T, b R) R {
	var r R

	return r
}

type M struct {
	Name string
}

type I interface {
	M
}

// error: interface contains type constraints
// func UseInterfaceEmbedStructDirect(i I) {

// }

func UseUseInterfaceEmbedStructGeneric[T I](t T) {
	// println(t.Name)

	// tt, ok := t.(M)
}
