package generic

import "testing"

// 返回零值
func Get[T any]() T {
	var t T
	return t
}

func Get2[T any]() (t T) {
	return
}

// 判断呢？
func IsZero[T any](t T) bool {
	// if t == 0 { // 可以吗?
	// if t == "" { // 可以吗?
	// if t == nil { // 可以吗?
	// if any(t) == nil { // 可以吗？
	// 如果都不可以，怎么样才行呢？
	//    return true
	// }

	// 用类型断言，拿到具体类型，再判断零值
	switch tt := any(t).(type) {
	case int:
		return tt == 0
	case string:
		return tt == ""

		// more ...
	}

	return false
}

func TestZero(t *testing.T) {
	t.Logf("Get: %+v\n", Get[int]())
	t.Logf("Get2: %q\n", Get2[string]())

	t.Logf("IsZero: %+v\n", IsZero(0))
	t.Logf("IsZero: %+v\n", IsZero(1))
	t.Logf("IsZero: %+v\n", IsZero(""))
	t.Logf("IsZero: %+v\n", IsZero("1"))
}
