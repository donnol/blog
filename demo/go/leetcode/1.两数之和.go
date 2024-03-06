package leetcode

var (
	_ = twoSum
)

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	l := len(nums)
	r := make([]int, 2)

	seen := make(map[int]int, l-1) // key is target-n, value is index
	for i := 0; i < l; i++ {
		n := nums[i]

		if ti, ok := seen[n]; ok {
			r[0] = ti
			r[1] = i
			break
		}

		if i == l-1 {
			break
		}

		seen[target-n] = i
	}

	return r
}

// @lc code=end
