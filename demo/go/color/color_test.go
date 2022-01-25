package color

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCar(t *testing.T) {
	// 超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 并发执行
	wg := new(sync.WaitGroup)

	for _, car := range []Runner{ // 遍历切片
		NewCar("lanbo", Red, 2),
		NewCar("boshi", Blue, 3),
	} {
		wg.Add(1) // 记录一个
		go func(car Runner) {
			defer wg.Done() // 完成一个

			t.Run(car.(*Car).name, func(t *testing.T) { // 对接口断言，获得具体类型
				car.Run(ctx)
			})
		}(car)
	}

	// 等上面均完成
	wg.Wait()

	timesMutex.RLock()
	fmt.Printf("times: %+v\n", times)
	timesMutex.RUnlock()
}
