package clonable

// Clonable defines a constraint of types having Clone() T method.
type Clonable[T any] interface {
	Clone() T
}

// 实现1
type m struct{}

func (m m) Clone() int {
	return 1
}

var _ Clonable[int] = m{}

// 实现2
type mc struct{}

func (m mc) Clone() mc {
	return mc{}
}

var _ Clonable[mc] = mc{}

// Repeat builds a slice with N copies of initial value.
// from github.com/samber/lo
// 关注类型约束：T Clonable[T]
// 其中T由Clonable约束，而Clonable接口接受T作为类型参数；
// 你中有我，我中有你 -- 互文
func Repeat[T Clonable[T]](count int, initial T) []T {
	result := make([]T, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}
	return result
}

// usage
var _ = func() int {
	// m类型的Clone方法返回的不是自身，所以不符合泛型约束：T Clonable[T]
	// m does not implement Clonable[m] (wrong type for method Clone)
	// 	have Clone() int
	// 	want Clone() m
	// Repeat[m](1, m{})

	// mc类型的Clone方法返回的恰好是自身，符合泛型约束
	Repeat(1, mc{})

	return 0
}
