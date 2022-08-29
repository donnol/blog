package executor

import (
	"fmt"
	"sync"
)

// 先后执行：一个操作的结果是另一个操作的输入
//
// 独立执行：多个操作互相独立没有先后依赖

type (
	Tie = any

	Executor[T Tie] interface {
		// 执行，按顺序执行每一批ops；同一批次可并发执行
		Exec(t T, bops ...[]Operation[T]) error
	}

	Operation[T Tie] func(T) error
)

func New[T Tie](
	concurrent bool,
) Executor[T] {
	return &executor[T]{
		concurrent: concurrent,
	}
}

type executor[T Tie] struct {
	concurrent bool // 是否并发
}

func (ex *executor[T]) Exec(t T, bops ...[]Operation[T]) error {
	if len(bops) == 0 {
		return fmt.Errorf("require at least one operation")
	}

	for _, ops := range bops {
		// 同批次，可并发执行
		if ex.concurrent {
			// 限制最大goroutine数
			limitCh := make(chan struct{}, 16)
			wg := new(sync.WaitGroup)
			errCh := make(chan error, 8)
			errs := make([]error, 0, 8)
			go func() {
				for err := range errCh {
					errs = append(errs, err)
				}
			}()
			for _, op := range ops {
				limitCh <- struct{}{}
				wg.Add(1)
				go func(op Operation[T]) {
					defer func() {
						<-limitCh
						wg.Done()
					}()

					if err := op(t); err != nil {
						errCh <- err
						return
					}
				}(op)
			}
			wg.Wait()
			close(limitCh)
			close(errCh)
			if len(errs) != 0 {
				return fmt.Errorf("exec failed: %+v", errs)
			}
		} else {
			for _, op := range ops {
				if err := op(t); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
