package leetcode

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
