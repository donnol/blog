package rotate

// Front 向前环转
func Front(s []int, index int) {
	FrontAny(s, index)
}

// Remove 移除指定位置的元素
func Remove(s *[]int, index int) {
	RemoveAny(s, index)
}

// 尝试使用以下函数签名时，发现[]int是不能直接传给[]any的，所以此时必须使用泛型
// func FrontAny1(s []any, index int) {
// }

func FrontAny[E any](s []E, index int) {
	if len(s) <= index {
		return
	}

	var e E
	for i := range s {
		switch {
		case i < index:
			continue
		case i >= index && i != len(s)-1:
			s[i] = s[i+1]
		case i == len(s)-1:
			s[i] = e
		}
	}
}

func RemoveAny[E any](s *[]E, index int) {
	if s == nil {
		return
	}
	if len(*s) <= index {
		return
	}

	FrontAny(*s, index)

	*s = (*s)[:len(*s)-1]
}
