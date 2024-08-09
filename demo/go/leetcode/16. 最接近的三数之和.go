package leetcode

import (
	"sort"
)

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。

// 示例 1：
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

// 示例 2：
// 输入：nums = [0,0,0], target = 1
// 输出：0

// 提示：
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	s := nums[0] + nums[1] + nums[2]
	d := abs(s - target)
	sum := s

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			s := nums[i] + nums[j] + nums[k]
			dt := abs(s - target)
			for j < k && dt > d {
				k--
				s = nums[i] + nums[j] + nums[k]
				dt = abs(s - target)
			}
			if j == k {
				break
			}
			if dt < d {
				d = dt
				sum = s
			}
		}
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
