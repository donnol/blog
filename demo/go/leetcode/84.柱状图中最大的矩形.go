package leetcode

var (
	_ = largestRectangleArea
)

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 1. 单体最大，h*1
	// 2. 连续最大，min(h1, h2, ...)*m*1，m为连续数量 -- 子集
	// 矩形面积由左边和右边柱子高度共同决定。如果左边柱子高度>当前柱子高度，那么还可以继续向左侧扩张；如果右边柱子高度>当前柱子高度，那么还可以继续向右侧扩张。所以我们要找两边下一个更矮的柱子来确定左右边界。这样就可以计算出以当前柱子高度为高，左右柱子间距为宽的矩形面积，并求出max

	l := len(heights)
	r := 0

	// 遍历的是高度，那么就寻找每个高度所能达到的最大面积--也就是如果两边有比它高的，那么它的宽度就能增大，从而增大它能达到的面积
	// 但是:
	// 	Time Limit Exceeded
	// 	91/98 cases passed (N/A)
	// TODO: 使用单调栈优化
	for i := 0; i < l; i++ {
		n := heights[i]

		li, ri := i, i
		for li > 0 && heights[li-1] >= n {
			li--
		}
		for ri < l-1 && heights[ri+1] >= n {
			ri++
		}

		max := n * (ri - li + 1)
		if max > r {
			r = max
		}
	}

	return r
}

// @lc code=end
