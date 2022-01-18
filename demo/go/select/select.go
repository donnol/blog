package main

import "fmt"

func main() {
	var a int

	// select会先执行每个case，如果其中有多个case有结果，会随机进入其中一个分支
	select {
	case n := <-f1(): // 每次执行都会输出：f1 run
		a = n
	case n := <-f2(): // 每次执行都会输出：f2 run
		a = n
	}

	fmt.Printf("a: %d\n", a) // 随机输出1或2
}

func f1() <-chan int {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	fmt.Printf("f1 run\n")

	return c
}

func f2() <-chan int {
	c := make(chan int)

	go func() {
		c <- 2
	}()

	fmt.Printf("f2 run\n")

	return c
}
