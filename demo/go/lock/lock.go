package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 从
// 一个CPU处理大量数据
// 到
// 多个CPU处理大量数据

var (
	// m *sync.Mutex // panic
	m sync.Mutex
	i int64
)

func UseLock() {
	// 一定要拿到锁，拿不到就执行不下去，所以只能不停地等
	m.Lock() // use atomic，如果多个goroutine都想拿锁，第一个成功后，后面的进入到饥饿状态，等待锁释放后获取
	i++
	m.Unlock()

	fmt.Printf("r: %d\n", i)
}

func Atomic() {
	n := atomic.AddInt64(&i, 1)

	fmt.Printf("r: %d\n", n)
}
