package leetcode

var (
	_ = setZeroes
)

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	// 最终目的是变为0，而依据则是某个点所在的十字(行或列)上是否存在0：存在则0，否则非0

	l := len(matrix)
	if l == 0 {
		panic("empty matrix")
	}
	w := len(matrix[0])

	// 只有一开始是0的才会有影响力，后面变0的没有
	// 怎么知道是不是刚赋值的0呢？

	// m*n
	// list := make([]int, l*w)

	// m+n
	// rows := make([]int, l)
	// cols := make([]int, w)

	// 使用矩阵本身的首行首列来存储每行每列是否有0
	// 为了区分原本的首行首列是否有0，需额外使用变量记录
	var r0, c0 = false, false
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			r0 = true
			break
		}
	}
	for i := range matrix {
		if matrix[i][0] == 0 {
			c0 = true
			break
		}
	}

	for i := 1; i < l; i++ {
		for j := 1; j < w; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < l; i++ {
		for j := 1; j < w; j++ {
			// 行或列: 存在则0，否则非0
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if r0 {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
	if c0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end
