package new_make

// Make init any type with new or make
func Make[T any](lencap ...int) T {
	var l, c int
	if len(lencap) >= 1 {
		l = lencap[0]
	}
	if len(lencap) >= 2 {
		c = lencap[1]
	}
	_, _ = l, c

	// https://github.com/golang/go/issues/45380#issuecomment-1158078224
	//
	// switch T { // 对类型T做switch
	// case AnySlice: // 代表任意slice的类型是什么呢？
	// 	return make(T, l, c)
	// case AnyMap: // 代表任意map的类型是什么呢？
	// 	return make(T, l)
	// case AnyChan: // 代表任意chan的类型是什么呢？
	// 	return make(T, l)
	// }

	tt := new(T)
	return *tt
}

// From https://github.com/golang/exp/blob/master/slices/slices.go
type AnySlice[T any] interface {
	~[]T
}

// From https://github.com/golang/exp/blob/master/maps/maps.go
type AnyMap[K comparable, V any] interface {
	~map[K]V
}

type AnyChan[T any] interface {
	~chan T
}
