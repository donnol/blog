package leetcode

import (
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

	// 有序
	sort.Ints(nums)

	// 双指针
	for i := 0; i < l; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] { // 排序之后，如果有连续相同的值，直接跳过
			continue
		}

		k := l - 1
		target := -1 * a

		for j := i + 1; j < l; j++ {
			if j > i+1 && nums[j] == nums[j-1] { // 排序之后，如果有连续相同的值，直接跳过
				continue
			}

			for j < k && nums[j]+nums[k] > target { // 比目标值大的情况下往前挪动`k`指针
				k--
			}
			if j == k { // 如果挪动到一样了还是大于目标值，说明找不到符合条件的值，直接跳出
				break
			}

			if nums[j]+nums[k] == target {
				r = append(r, []int{a, nums[j], nums[k]})
			}
		}

	}
	return r
}
