package struct_method_and_func_field

type M struct {
	// 函数类型字段，无法获取M实例
	// 如若需要M实例，需要将函数签名改为`func(*M)`
	// f func()
	f func(*M)

	f1 func(M)
}

// 方法则不一样，可以直接获得M的实例：m
func (m *M) Method() {
	m.f(m)

	m.f1(*m)
}

// === 发散 ===

// 如果结构体定义支持Self特殊结构，表示结构体自身实例
type (
	Self = any // 当成为语言特性后，不需要此定义

	M1 struct {
		// 此时，函数参数类型是*M1
		F func(*Self)
	}
)

type M2 struct {
	Clone func(*M2) *M2
	Copy  func(M2) M2

	// illegal cycle in declaration of M2
	// other M2

	OtherPtr *M2 // 使用指针可以

	Other Self // 但只是想使用值呢？
}
