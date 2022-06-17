package new_make

import (
	"testing"
	"unsafe"
)

func BenchmarkNewEmptyStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := new(struct{})
		_ = a
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := new(A)
		_ = a
	}
}

func TestEmptyStruct(t *testing.T) {
	s := struct{}{}
	_ = s
	t.Logf("size: %v\n", unsafe.Sizeof(s))

	a := new(struct{})
	_ = a
	t.Logf("size: %v, %v\n", unsafe.Sizeof(a), unsafe.Sizeof(*a))
}

func TestStruct(t *testing.T) {
	s := A{}
	_ = s
	t.Logf("size: %v\n", unsafe.Sizeof(s))

	a := new(A)
	_ = a
	t.Logf("size: %v, %v\n", unsafe.Sizeof(a), unsafe.Sizeof(*a))
}

func TestDerefer(t *testing.T) {
	var a *A
	_ = a
	// t.Logf("a: %v\n", *a) // panic
	// t.Logf("a.String: %q\n", a.String()) // panic

	var b = new(A)
	t.Logf("b: %v\n", *b)                // work
	t.Logf("b.String: %q\n", b.String()) // work
}
