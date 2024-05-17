package leetcode

import "fmt"

// 每增加一个数组都有可能改变之前的交集
func intersectionSizeTwo(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	curr := intervals[0]
	nums := make([]int, 0, len(intervals))
	for i := 1; i < len(intervals); i++ {
		// 本组数字与下一组数字取交集
		next := intervals[i]
		inter := intersect(curr, next, 2)
		nums = append(nums, inter...)

		curr = next
	}

	// 记录数字出现次数，不小于2表示该数字为交集
	// 若交集数字够数，则返回
	// 若交集数字不够，则从各组数字取非交集数字，直到够数
	// 但数量并不确定，怎么知道何时够数呢？

	seen := make(map[int]int) // 数字出现的次数
	for i := 0; i < len(intervals); i++ {
		for j := intervals[i][0]; j <= intervals[i][1]; j++ {
			if seen[j] >= 1 {
				nums = append(nums, j)
			}
			seen[j]++
		}
	}
	fmt.Println(seen)

	return len(nums)
}

// curr, next均是两个元素，其表示从元素0到元素1的递增数组，[1, 3] -> [1, 2, 3]
//
// 有序数组，使用双指针遍历比较
//
// 这样子只会拿到前n个交集元素
func intersect(curr, next []int, n int) []int {
	r := make([]int, 0, n)

	cindex, nindex := 0, 0
	for cindex < (curr[1]-curr[0]+1) && nindex < (next[1]-next[0]+1) {
		cnum := cindex + curr[0]
		nnum := nindex + next[0]
		if cnum == nnum {
			r = append(r, cnum)
			if len(r) >= n {
				break
			}

			cindex++
			nindex++
		} else if cnum < nnum {
			cindex++
		} else {
			nindex++
		}
	}

	return r
}
