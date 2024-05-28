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
