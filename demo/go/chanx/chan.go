package chanx

import "fmt"

// oneSendMultiRecv
//
// 一个无缓冲管道，有一个发送者，两个接收者
// 谁会接收到呢？
// 后进先出，2接收到？
func oneSendMultiRecv() {
	ch := make(chan int)

	go func() {
		i := <-ch
		fmt.Printf("1 i: %d\n", i)
	}()

	go func() {
		i := <-ch
		fmt.Printf("2 i: %d\n", i)
	}()

	ch <- 1
}

func oneSendManyMultiRecv() {
	ch := make(chan int)

	go func() {
		i := <-ch
		fmt.Printf("1 i: %d\n", i)
	}()

	go func() {
		i := <-ch
		fmt.Printf("2 i: %d\n", i)
	}()

	ch <- 1
	ch <- 2
}

func multiSendMultiRecv() {
	ch := make(chan int)

	go func() {
		i := <-ch
		fmt.Printf("1 i: %d\n", i)
	}()

	go func() {
		i := <-ch
		fmt.Printf("2 i: %d\n", i)
	}()

	wait := make(chan bool)
	go func() {
		ch <- 1

		wait <- true
	}()

	ch <- 2

	<-wait
}
