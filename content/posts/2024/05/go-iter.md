# Go range over func

[Play](https://go.dev/play/p/BR1sVgX_ZSn?v=gotip)

```go
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"iter"
)

func count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				break
			}
		}
	}
}

// Tree is a binary tree.
type Tree[E any] struct {
	val         E
	left, right *Tree[E]
}

// All may be used in a for/range loop to iterate
// through all the values of the tree.
// This implementation does an in-order traversal.
func (t *Tree[E]) All(yield func(E) bool) {
	t.doAll(yield)
}

// doAll is a helper for All, to make it easier
// to know when the iteration stopped in a subtree.
func (t *Tree[E]) doAll(yield func(E) bool) bool {
	if t == nil {
		return true
	}
	return t.left.doAll(yield) && yield(t.val) && t.right.doAll(yield)
}

func SumTree(t *Tree[int]) int {
	s := 0
	for v := range t.All {
		s += v
	}
	return s
}

func SumTreeSeq(seq iter.Seq[int]) int {
	s := 0
	for v := range seq {
		s += v
	}
	return s
}

func main() {
	fmt.Println("Hello, 世界")

	// 函数body会被转为yield函数传入到iter.Seq里 -- 当`body`没有控制流语句时，一律视为`return true`
	for k := range count(10) {
		fmt.Println(k)
	}

	t := &Tree[int]{
		val:   2,
		left:  &Tree[int]{val: 1},
		right: &Tree[int]{val: 3},
	}
	r := SumTree(t)
	fmt.Println(r, SumTreeSeq(t.All))
}

```
