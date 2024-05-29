package leetcode

func MaxLen(s []string) int {
	return maxLen(s)
}

func maxLen(s []string) int {
	max := 0
	for i := range s[:len(s)-1] {
		curr := s[i]
		next := s[i+1]

		if haveSameChar(curr, next) {
			continue
		}

		r := len(curr) * len(next)
		if r > max {
			max = r
		}
	}
	return max
}

func haveSameChar(s1, s2 string) bool {
	seen := make(map[byte]struct{})
	for i := range s1 {
		seen[s1[i]] = struct{}{}
	}
	for i := range s2 {
		_, ok := seen[s2[i]]
		if ok {
			return true
		}
	}
	return false
}

func MaxProduct(s []string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	masks := make([]int64, n)
	for i := range s {
		for j := range s[i] {
			masks[i] |= 1 << (s[i][j] - 'a')
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i]&masks[j] == 0 {
				dp[i][j] = len(s[i]) * len(s[j])
			}
			ai, aj := i-1, j-1
			if ai < 0 {
				ai = 0
			}
			if aj < 0 {
				aj = 0
			}
			dp[i][j] = max(dp[i][j], dp[ai][j], dp[i][aj])
		}
	}

	return dp[n-2][n-1]
}

func max(s ...int) int {
	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

// def maxProduct(words):
//     n = len(words)
//     masks = [0] * n  # 存储每个字符串的位掩码
//     dp = [[0] * n for _ in range(n)]  # 动态规划表

//     # 计算每个字符串的位掩码
//     for i, word in enumerate(words):
//         for c in word:
//             masks[i] |= 1 << (ord(c) - ord('a'))

//     # 动态规划
//     for i in range(n):
//         for j in range(i + 1, n):
//             if masks[i] & masks[j] == 0:  # 检查位掩码是否有交集
//                 dp[i][j] = len(words[i]) * len(words[j])
//             dp[i][j] = max(dp[i][j], dp[i-1][j], dp[i][j-1])  # 更新最大值

//     return dp[n - 2][n - 1]

// # 示例
// words = ["abcw", "baz", "foo", "bar", "xtfn", "abcdef"]
// result = maxProduct(words)
// print(result)  # 输出：16
