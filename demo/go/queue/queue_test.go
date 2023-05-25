package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()

	queue.Push("hello")
	e := queue.Pop()
	if e != "hello" {
		t.Errorf("bad pop, %s != %s", e, "hello")
	}

	// listen
	queue.Listen(func(e string) {
		fmt.Printf("got ele: %s\n", e)
	})

	queue.Push("hello")
	time.Sleep(time.Second)

	queue.Push("world")
}
