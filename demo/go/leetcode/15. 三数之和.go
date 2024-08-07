package leetcode

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。

// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。

// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0 。

// 提示：
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105
func threeSum(nums []int) [][]int {
	l := len(nums)
	r := make([][]int, 0, l/2)
	seen := make(map[string]struct{})
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				a := nums[i]
				b := nums[j]
				c := nums[k]

				if a+b+c == 0 {
					s := []int{a, b, c}
					sort.Ints(s)
					key := fmt.Sprintf("%d|%d|%d", s[0], s[1], s[2])
					_, ok := seen[key]
					if !ok {
						r = append(r, s)
						seen[key] = struct{}{}
					}
				}
			}
		}

	}
	return r
}
