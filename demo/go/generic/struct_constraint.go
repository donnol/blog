package generic

import (
	"fmt"
)

// https://github.com/golang/go/issues/50417

type Condition[T ~string, S ~string] interface {
	~struct{
		Type               T
		Status             string
		Severity           S
		LastTransitionTime int64
		Reason             string
		Message            string
	}
}

// structField is a type constraint whose type set consists of some
// struct types that all have a field named x.
type structField interface {
	struct { a int; x int } |
		struct { b int; x float64 } |
		struct { c int; x uint64 }
}

// This function is INVALID.
// func IncrementX[T structField](p *T) {
// 	v := p.x // INVALID: p.x undefined (type *T has no field or method x)
// 	v++
// 	p.x = v // INVALID: p.x undefined (type *T has no field or method x)
// }

type structFieldSameX interface {
	struct { a int; X int } |
		struct { b int; X int } |
		struct { c int; X int }
}

// 同样类型的x字段，还是不行呢: go version devel go1.18-07525e1 Fri Jan 7 00:15:59 2022 +0000 linux/amd64
// func IncrementSameX[T structFieldSameX](p *T) {
// 	v := p.X 
// 	v++
// 	p.X = v 
// }

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

type Person struct {
	Name string
}

type Named interface {
	~struct{ Name string }
}

func SayName[T Named](t T) {
	fmt.Println(t.Name)
}

type NamedC interface {
	// 如果不注释掉，会报错：(SayNameC: t.Name) t.Name undefined (interface NamedC has no method Name)
	// struct {a int; Name string} |
	// struct {b int; Name string} |
	struct {c int; Name string} 
}

func SayNameC[T NamedC](t T) {
	fmt.Println(t.Name)
}
