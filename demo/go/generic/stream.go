package generic

// 目前不支持方法泛型，所以需要在结构体定义里把结果类型也一并声明；
// 事实上，返回的结果应该是由后续调用的方法来定的
type Stream[T, R any] struct {
	isHead bool
	next   *Stream[T, R]
	f      func(T) R
	data   []T
}

func New[T, R any](data []T) *Stream[T, R] {
	return &Stream[T, R]{
		isHead: true,
		data:   data,
	}
}

func (stream *Stream[T, R]) Map(f func(T) R) *Stream[T, R] {

	return stream
}

func (stream *Stream[T, R]) Reduce(f func(T) R) *Stream[T, R] {

	return stream
}

func (stream *Stream[T, R]) Exec() []R {
	var r []R

	return r
}

func (stream *Stream[T, R]) ForEach(f func(T)) {

}

type Stream2[T any] struct {
	data []T
}

func New2[T any](data []T) *Stream2[T] {
	return &Stream2[T]{
		data: data,
	}
}

// 不支持在方法里另外声明类型参数
// func (stream *Stream2[T]) Map[R any](f func(T) R) *Stream2[T] {

// 	return stream
// }
