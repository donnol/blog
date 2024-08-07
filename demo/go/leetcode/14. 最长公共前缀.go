package leetcode

import "strings"

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"

// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

// 提示：

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
func longestCommonPrefix(strs []string) string {
	short := strs[0]
	for _, s := range strs {
		if len(s) < len(short) {
			short = s
		}
	}
	others := []string{}
	for _, s := range strs {
		if s == short {
			continue
		}
		others = append(others, s)
	}

	r := ""
	for j := 1; j <= len(short); j++ {
		s := short[0:j]

		exist := true
		for _, other := range others {
			if !strings.HasPrefix(other, s) {
				exist = false
			}
		}
		if exist && len(s) > len(r) {
			r = s
		}
	}
	return r
}
