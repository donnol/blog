package str

import "fmt"

func ToInt(s string) int {
	n := len(s)
	r := 0

	if n <= 0 {
		panic(fmt.Errorf("input is empty"))
	}

	var i int
	var sign = 1
	if s[i] == '-' {
		sign = -1
		i++
	}

	for ; i < n; i++ {
		if s[i] > '9' || s[i] < '0' {
			panic(fmt.Errorf("invalid char %c", s[i]))
		}

		digit := s[i] - '0'
		r = r*10 + int(digit)
	}

	r *= sign

	return r
}
