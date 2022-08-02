package convert

import (
	"reflect"
	"testing"
)

type UserTable struct {
	Id   uint
	Name string
}

type UserParam struct {
	Name string
}

// ByName return a ConvertFunc which use field name as map condition
func ByName[T, R any]() ConvertFunc[T, R] {
	return func(t T) R {
		var r R

		// 反射

		// 生成

		return r
	}
}

// 这个方法如果还是要手写，也挺麻烦的；特别是字段很多的时候
// 如果可以根据约束里的结构体来生成，就方便多了
// 不过，生成的时候如何对应呢？如果都是以字段名，当然是方便，但如果是以其它呢？
func UserParamToTableByName[T UserParam, R UserTable]() ConvertFunc[T, R] {
	return func(t T) R {
		var r R

		// 明明都能用结构体来做约束了，但是却还是没法用它的字段
		// https://github.com/golang/go/issues/48522
		// r.Name = t.Name // not support in go1.18 and go1.19 will not too

		return r
	}
}

func TestUserParamToTableByName(t *testing.T) {
	in := []UserParam{
		{"jd"},
		{"je"},
	}
	r := Convert(in, UserParamToTableByName())
	t.Logf("r: %v\n", r)
}

func TestItoa(t *testing.T) {
	in := []int{
		1, 2, 3,
	}
	r := Convert(in, Itoa2[int])
	want := []string{"1", "2", "3"}
	if !reflect.DeepEqual(r, want) {
		t.Fatalf("bad case, %v != %v", r, want)
	}
}
