package str

// Index 字符串匹配的KMP算法实现(O(n+m))
func Index(s, pattern string) int {
	nex := next(pattern)

	tar, pos := 0, 0

	for tar < len(s) {
		if s[tar] == pattern[pos] {
			tar++
			pos++
		} else if pos != 0 {
			pos = nex[pos-1]
		} else {
			tar++
		}

		if pos == len(pattern) {
			return tar - pos
		}
	}

	return -1
}

// next 为子串建立next数组，在遍历过程中跳过不可能匹配上的内容，从而减少遍历次数
// "部分匹配值"就是"前缀"和"后缀"的最长的共有元素的长度
// 前缀: 指一个字符串从开头到某个位置结束的子串，不包含最后一个字符。
// 后缀: 指一个字符串从某个位置开始到结尾的子串，不包含第一个字符。
// next数组元素：代表当前字符之前的字符串中，"前缀"和"后缀"的最长的共有元素的长度
// 可以确切地知道在当前位置之前的一个潜在匹配的位置
func next(pattern string) []int {
	r := make([]int, 0, len(pattern))

	r = append(r, 0)
	x := 1
	now := 0

	for x < len(pattern) {
		if pattern[now] == pattern[x] {
			now += 1
			x += 1
			r = append(r, now)
		} else if now != 0 {
			now = r[now-1]
		} else {
			r = append(r, 0)
			x += 1
		}
	}

	return r
}
