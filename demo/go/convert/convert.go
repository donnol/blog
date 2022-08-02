package convert

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type ConvertFunc[T, R any] func(T) R

func Convert[T, R any](in []T, f ConvertFunc[T, R]) []R {
	r := make([]R, len(in))
	for i := 0; i < len(in); i++ {
		r[i] = f(in[i])
	}
	return r
}

func ConvertOne[T, R any](in T, f ConvertFunc[T, R]) R {
	return Convert([]T{in}, f)[0]
}

// Itoa return a ConvertFunc which convert int to string
func Itoa[T constraints.Integer, R string]() ConvertFunc[T, R] {
	return func(t T) R {
		return R(strconv.Itoa(int(t)))
	}
}

func Itoa2[T constraints.Integer, R string](t T) R {
	return R(strconv.Itoa(int(t)))
}
