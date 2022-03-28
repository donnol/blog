package lo

import (
	"testing"

	"github.com/samber/lo"
)

func TestLo(t *testing.T) {
	r := lo.Map([]int{1, 2, 3}, func(t int, i int) int {
		return t + i // i是遍历时的下标，t是元素
	})
	t.Logf("r: %+v\n", r) // [1, 3, 5]

	re := lo.Reduce([]int{1, 2, 3}, func(r int, t int, i int) int {
		return r + t // r是初始值（也就是Reduce的最后一个参数），t是元素值，i是下标
	}, 0)
	t.Logf("re: %+v\n", re)

	// 级联
	re = lo.Reduce(lo.Map([]int{1, 2, 3}, func(t int, i int) int {
		return t + i // i是遍历时的下标，t是元素
	}), func(r int, t int, i int) int {
		return r + t
	}, 0)
	t.Logf("re: %+v\n", re)
}
