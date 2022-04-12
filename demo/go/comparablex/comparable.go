package comparablex

import (
	"go/types"
	"reflect"
)

type Set[E comparable] []E // 可以用做类型参数的约束

// 使用go1.18编译，报错：interface is (or embeds) comparable
// var A comparable // 变量不可以使用`comparable`类型

// 同样是内置的error类型，就能给变量使用
var E error

func ReflectComparable(v interface{}) bool {
	typ := reflect.TypeOf(v)

	// Comparable reports whether values of this type are comparable.
	// Even if Comparable returns true, the comparison may still panic.
	// For example, values of interface type are comparable,
	// but the comparison will panic if their dynamic type is not comparable.
	// -- 即使返回true，也有可能panic。
	// 比如：接口类型的值是可比较的，但如果它们的动态类型是不可比较的，就会panic
	return typ.Comparable()
}

func TypesComparable() bool {
	t := types.NewChan(types.SendOnly, &types.Basic{})

	return types.Comparable(t)
}
