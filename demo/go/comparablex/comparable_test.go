package comparablex

import (
	"reflect"
	"testing"

	"github.com/donnol/blog/demo/go/comparablex/inner"
)

type I interface {
	Name() string
}

type II1 struct {
	name string
}

func (i II1) Name() string {
	return i.name
}

type II2 struct {
	name string

	m map[int]string
}

func (i II2) Name() string {
	return i.name
}

type Slice interface {
	Len() int
}

type sliceImpl []int

func (impl sliceImpl) Len() int {
	return len(impl)
}

func TestReflectComparable(t *testing.T) {
	type args struct {
		v interface{}
	}

	var (
		p  *int
		ch chan int
		s  struct {
			id uint
		}
		arr [2]int

		sli []int
		m   map[int]string
		f   func()

		// i I // `typ.Comparable()`会panic
		i1  I = (*II1)(nil)
		i2  I = (*II2)(nil)
		ii1 I = II1{}
		ii2 I = II2{}

		si Slice = sliceImpl{}

		reftype = reflect.TypeOf(inner.Type{})

		ei1 inner.EI = inner.EI1{}
	)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{v: 1}, true},
		{"", args{v: 1.0}, true},
		{"", args{v: "1"}, true},
		{"", args{v: true}, true},
		{"", args{v: false}, true},
		{"", args{v: p}, true},
		{"", args{v: ch}, true},
		{"", args{v: s}, true},
		{"", args{v: arr}, true},
		{"", args{v: i1}, true},
		{"", args{v: i2}, true},
		{"", args{v: ii1}, true},
		{"", args{v: reftype}, true},
		{"", args{v: ei1}, true},

		{"", args{v: sli}, false},
		{"", args{v: m}, false},
		{"", args{v: f}, false},
		{"", args{v: ii2}, false},
		{"", args{v: si}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReflectComparable(tt.args.v); got != tt.want {
				t.Errorf("ReflectComparable() = %v, want %v", got, tt.want)
			}
		})
	}

	var ii11 I = II1{}
	t.Logf("ii1 = ii11: %t", ii1 == ii11)

	// panic: runtime error: comparing uncomparable type comparablex.II2 [recovered]
	// panic: runtime error: comparing uncomparable type comparablex.II2
	// var ii22 I = II2{}
	// if ii2 == ii22 {
	// 	t.Logf("ii2 = ii22")
	// }

	var (
		i11 I = (*II1)(nil)
		i22 I = (*II2)(nil)
	)
	t.Logf("i1 = i11: %t", i1 == i11)
	t.Logf("i2 = i22: %t", i2 == i22)

	var (
		reftype2 = reflect.TypeOf(inner.Type{})
	)
	t.Logf("reftype = reftype2: %t", reftype == reftype2)

	var (
		ei2 inner.EI = inner.EI1{}
	)
	t.Logf("ei1 = ei2: %t", ei1 == ei2)

	// 返回true，会panic的情况
	// 怎么触发这种情况呢？
}

func TestTypesComparable(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypesComparable(); got != tt.want {
				t.Errorf("TypesComparable() = %v, want %v", got, tt.want)
			}
		})
	}
}
