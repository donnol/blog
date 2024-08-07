package leetcode

// 七个不同的符号代表罗马数字，其值如下：

// 符号	值
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// 罗马数字是通过添加从最高到最低的小数位值的转换而形成的。将小数位值转换为罗马数字有以下规则：

// 如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
// 如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号，例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX。仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)。
// 只有 10 的次方（I, X, C, M）最多可以连续附加 3 次以代表 10 的倍数。你不能多次附加 5 (V)，50 (L) 或 500 (D)。如果需要将符号附加4次，请使用 减法形式。
// 给定一个整数，将其转换为罗马数字。

// 示例 1：
// 输入：num = 3749
// 输出： "MMMDCCXLIX"
// 解释：
// 3000 = MMM 由于 1000 (M) + 1000 (M) + 1000 (M)
//  700 = DCC 由于 500 (D) + 100 (C) + 100 (C)
//   40 = XL 由于 50 (L) 减 10 (X)
//    9 = IX 由于 10 (X) 减 1 (I)
// 注意：49 不是 50 (L) 减 1 (I) 因为转换是基于小数位

// 示例 2：
// 输入：num = 58
// 输出："LVIII"
// 解释：
// 50 = L
//
//	8 = VIII
//
// 示例 3：
// 输入：num = 1994
// 输出："MCMXCIV"
// 解释：
// 1000 = M
//
//	900 = CM
//	 90 = XC
//	  4 = IV
//
// 1 <= num <= 3999
func intToRoman(num int) string {
	r := ""

	i := 0
	for {
		mod := num % 10
		p10 := binPow(10, i)
		e := mod * p10
		big, mid, small := baseByNum(e)
		switch mod {
		case 4, 9:
			r = small + big + r
		default:
			if mid != "" {
				r = mid + repeat(small, times(mod)) + r
			} else {
				c := big
				if mod < 5 {
					c = small
				}
				r = repeat(c, times(mod)) + r
			}
		}

		i++
		num /= 10
		if num == 0 {
			break
		}
	}

	return r
}

// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
func baseByNum(num int) (big, mid, small string) {
	switch {
	case num == 1:
		big = "I"
		small = "I"
	case num < 5:
		big = "V"
		small = "I"
	case num == 5:
		big = "V"
		small = "V"
	case num < 10:
		big = "X"
		mid = "V"
		small = "I"
	case num == 10:
		big = "X"
		small = "X"
	case num < 50:
		big = "L"
		small = "X"
	case num == 50:
		big = "L"
		small = "L"
	case num < 100:
		big = "C"
		mid = "L"
		small = "X"
	case num == 100:
		big = "C"
		small = "C"
	case num < 500:
		big = "D"
		small = "C"
	case num == 500:
		big = "D"
		small = "D"
	case num < 1000:
		big = "M"
		mid = "D"
		small = "C"
	default:
		big = "M"
		small = "M"
	}
	return
}

func binPow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 != 0 {
			res = res * a
		}
		a = a * a
		b >>= 1
	}
	return res
}

func repeat(s string, n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r += s
	}
	return r
}

func times(mod int) int {
	n := 0
	if mod == 5 {
		return 1
	}
	if mod > 0 && mod < 4 {
		n = mod
	}
	if mod > 5 && mod < 9 {
		n = mod - 5
	}
	return n
}
