package sched

import (
	"fmt"
	"runtime"
)

var (
	_ = runtime.Version()
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// 如果没有这个调用，主线程的say先执行完，才轮到次线程
		// 如果添加这个调用，线程会主动让出CPU，另外的线程(可能)会得到先执行的机会
		runtime.Gosched()
		fmt.Println(s)
	}
}

func Demo() {
	runtime.GOMAXPROCS(1)

	go say("world")
	say("hello")
}
