package stack

import (
	"math/rand"
	"time"

	"github.com/donnol/blog/demo/go/fmtx"
)

func NormalStack(arr []int) {
	n := len(arr)
	stack := make([]int, 0, n)

	// push
	for i := 0; i < n; i++ {
		stack = append(stack, arr[i])

		// top
		top := stack[len(stack)-1]
		fmtx.Printf("top: %d\n", top)
	}

	// pop
	for i := 0; i < n; i++ {
		index := len(stack) - 1
		out := stack[index] // last one
		fmtx.Printf("out: %d\n", out)

		stack = stack[:index]
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// mono increase from bottom to top
func MonoIncrStack(arr []int) {
	n := len(arr)
	stack := make([]int, n)

	for i := 0; i < n; i++ {
		stack[i] = arr[i]
	}

	fmtx.Printf("before stack: %+v\n", stack)

	// push
	// insert sort to handle array
	// 原地排序，减少内存分配
	for i := 1; i < n; i++ {
		temp := stack[i]
		j := i - 1
		for ; j >= 0; j-- {
			if stack[j] <= temp {
				break
			}
			stack[j+1] = stack[j]
		}
		stack[j+1] = temp
	}
	fmtx.Printf("after  stack: %+v\n", stack)

	// insert one
	var last = stack[n-1]
	var si = 50
	for i := 1; i < n; i++ {
		if si >= stack[i] {
			continue
		}

		for j := n - 1; j > i; j-- {
			stack[j] = stack[j-1]
		}
		stack[i] = si
		break
	}
	stack = append(stack, last)

	fmtx.Printf("after  stack: %+v\n", stack)

	// var si = 100
	// pop
	// 这样写会由于循环内修改了stack(stack = stack[:index])导致循环次数减半
	// for i := 0; i < len(stack); i++ {
	// 因此，改用l变量，在循环前先拿一次长度
	l := len(stack)
	for i := 0; i < l; i++ {
		index := len(stack) - 1
		out := stack[index] // last one
		_ = out
		fmtx.Printf("out: %d\n", out)

		stack = stack[:index]
	}
}

// mono decrease from bottom to top
func MonoDecrStack(arr []int) {

}
