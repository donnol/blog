package pipe

type Pipe[T any] struct {
	e T

	fs []func()
}

func NewPipe[T any](fs ...func()) *Pipe[T] {
	return &Pipe[T]{
		fs: fs,
	}
}

func (p *Pipe[T]) Do() {
	// 怎么让fs装入各种类型的函数
	// 怎么将e传入函数
	_ = p.e
	for _, fun := range p.fs {
		fun()
	}
}

// 有限数量，格式一致的呢？
