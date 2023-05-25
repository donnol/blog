package queue

import "sync"

type Queue[T any] struct {
	mu   sync.Mutex
	ch   chan struct{} // 当有数据到来时通知
	eles []T           // 实际元素存储
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		ch:   make(chan struct{}, 1),
		eles: make([]T, 0, 32),
	}
}

func (queue *Queue[T]) Push(e T) {
	queue.mu.Lock()
	queue.eles = append(queue.eles, e)
	queue.mu.Unlock()

	queue.ch <- struct{}{}
}

func (queue *Queue[T]) Pop() (e T) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	if len(queue.eles) == 0 {
		return
	}
	e = queue.eles[0]
	queue.eles = queue.eles[1:]

	// 还有元素的情况下，因为这次pop把ch消费了，所以需要重新再发一次，使得后续的元素可以被监听器获取到
	if len(queue.eles) > 0 {
		queue.ch <- struct{}{}
	}

	return
}

func (queue *Queue[T]) Listen(f func(e T)) {
	go func() {
		for range queue.ch {
			e := queue.Pop()
			f(e)
		}
	}()
}
