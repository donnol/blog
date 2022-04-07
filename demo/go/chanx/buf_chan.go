package chanx

import (
	"fmt"
	"time"
)

// 一个带缓存的chan，是在它满了才开始消费，还是在它刚有数据就开始消费呢？
//
// 有数据就开始消费了。

func bufChan() {
	ch := make(chan int, 10)

	// 消费
	go func() {
		for e := range ch {
			fmt.Printf("e: %d\n", e)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 生产
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
