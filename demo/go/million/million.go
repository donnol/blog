package million

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// start 当k越来越大时，使用的内存也越来越大；调度耗费的时间也会相应提高；因此，在实际使用中，最好限制goroutine的数量，个人建议在`10k`以内；如果实际执行的任务超过该数字，则排队逐个处理。
func start(numRoutines int) {
	var wg sync.WaitGroup
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(10 * time.Second)

			if i == numRoutines-1 {
				ms := &runtime.MemStats{}
				runtime.ReadMemStats(ms)
				log.Printf("memory stats: %v, %v, %v\n", ms.Alloc, ms.HeapInuse, ms.StackInuse)
			}
		}(i)
	}
	wg.Wait()
}

// === RUN   Test_start/1k
// 2023/05/23 15:22:16 memory stats: 787608, 1835008, 8781824
// === RUN   Test_start/10k
// 2023/05/23 15:22:26 memory stats: 5879680, 6930432, 82837504
// === RUN   Test_start/100k
// 2023/05/23 15:22:36 memory stats: 56368464, 58605568, 821526528
// === RUN   Test_start/1m
// 2023/05/23 15:22:49 memory stats: 559107344, 575168512, 8192917504
