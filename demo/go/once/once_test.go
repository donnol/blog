package once

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	once.Do(func() {
		go func() {
			fmt.Println("start")
			time.Sleep(1 * time.Second)
			fmt.Println("stop")
		}()
	})

	time.Sleep(2 * time.Second)
}
