package leetcode

import "fmt"

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
// - 0 1 2 3 4 5 6
// 0 P   A   H   N  -- 坐标差值：P(0)      A(4)      H(8)       N(12) 即 3*2-2=4
// 1 A P L S I I G  -- 坐标差值：A(1) P(3) L(5) S(7) I(9) I(11) G(13) 即 2
// 2 Y   I   R      -- 坐标差值：Y(2)      I(6)      R(10)            即 3*2-2=4
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
// 请你实现这个将字符串进行指定行数变换的函数：
// string convert(string s, int numRows);
// 示例 1：
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// - 0 1 2 3 4 5 6
// 0 P     I    N  -- 坐标差值：P(0)        I(6)         N(12) 即 4*2-2=6
// 1 A   L S  I G  -- 坐标差值：A(1)   L(5) S(7)   I(11) G(13) 即 4,2,4,2
// 2 Y A   H R     -- 坐标差值：Y(2) A(4)   H(8) R(10)         即 2,4,2
// 3 P     I       -- 坐标差值：P(3)        I(9)               即 4*2-2=6
// 示例 3：
// 输入：s = "A", numRows = 1
// 输出："A"

// s = "PAYPALISHIRING", numRows = 5
// - 0 1 2 3 4 5 6
// 0 P       H     -- 坐标差值：P(0)          H(8)         即 5*2-2=8
// 1 A     S I     -- 坐标差值：A(1)     S(7) I(9)         即 6,2
// 2 Y   I   R     -- 坐标差值：Y(2)   I(6)   R(10)        即 4,4
// 3 P L     I G   -- 坐标差值：P(3) L(5)     I(11) G(13)  即 2,6,2
// 4 A       N     -- 坐标差值：A(4)          N(12)        即 5*2-2=8
// 提示：
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
func convert(s string, numRows int) string {
	r := ""

	l := len(s)
	equo := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		equo[i] = make([]byte, 0, l/2)
	}
	// defer func() {
	// 	print(equo)
	// }()

	rn := 0
	cn := 0
	j := 0
	for i := 0; i < l; i++ {
		e := s[i]

		if cn == 0 {
			equo[rn] = append(equo[rn], e)
			rn++
			if rn >= numRows {
				j++
				rn = 0

				cn++
				if cn >= numRows-1 {
					cn = 0
				}
			}
		} else {
			if rn == 0 {
				rn = numRows - 2
			}
			equo[rn] = append(equo[rn], e)
			j++

			rn--

			cn++
			if cn >= numRows-1 {
				cn = 0
			}
		}
	}

	for _, item := range equo {
		es := make([]byte, 0, len(item))
		for _, e := range item {
			if e != 0 {
				es = append(es, e)
			}
		}
		r += string(es)
	}
	return r
}

var _ = print

func print(equo [][]byte) {
	for _, item := range equo {
		esz := make([]byte, 0, len(item))
		esz = append(esz, item...)
		fmt.Println(string(esz))
	}
}
