package leetcode

import (
	"fmt"
	"strconv"
)

// 给定一个输入 n ，输出一个包含 n 个元素数组，数组中每个元素表示其索引的二进制形式中的数字 1 的个数
//
// 0,  1,  2,  3,   4,   5,   6,   7,    8,    9,   10,   11,   12,   13,   14,   15,    16,    17
// 0, 01, 10, 11, 100, 101, 110, 111, 1000, 1001, 1010, 1011, 1100, 1101, 1110, 1111, 10000, 10001
// 0,  1,  1,  2,   1,   2,   2,   3,    1,    2,    2,    3,    2,    3,    3,    4,     1,     2
//
// 数字是连续的，那么，可不可以根据上一个数字的1个数推出下一个数的1个数呢？如果可以，推导规则是什么呢？
//
// 2的整次幂时，1个数开始变小为1，然后再渐增
// 渐增规则：(i)/last(2^n) + 1 + if(i%last(2^(n-1))>2, 1, 0)

func Solution(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = amountOf1(i)
	}
	return r
}

// 数字i的二进制形式中1的个数
func amountOf1(i int) (r int) {
	s := fmt.Sprintf("%b", i)
	for _, b := range []byte(s) {
		if b == '1' {
			r++
		}
	}
	return
}

func Solution2(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = amountOf1V2(i)
	}
	return r
}

// 数字i的二进制形式中1的个数
func amountOf1V2(i int) (r int) {
	s := strconv.FormatInt(int64(i), 2)
	for _, b := range s {
		if b == '1' {
			r++
		}
	}
	return
}

func Solution3(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		sr := 0

		// (1)与0b1位与，当结果不为0，则表明数字的最后一位是1
		// (2)再把数字向右移一位，继续(1)，直到数字为0
		ii := i
		for ii != 0 {
			if ii&0b1 != 0 {
				sr++
			}
			ii >>= 1
		}

		r[i] = sr
	}
	return r
}

var _ = amountOf1V3

func amountOf1V3(i int) (r int) {
	for i != 0 {
		if i&0b1 != 0 {
			r++
		}
		i >>= 1
	}

	return
}

func Solution4(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		if i <= 2 {
			r[i] = 1
			continue
		}

		r[i] = amountOf1V4(uint32(i))
	}
	return r
}

var _ = amountOf1V4

func amountOf1V4(i uint32) (r int) {
	for i != 0 {
		i &= i - 1
		r++
	}
	return
}

func IsPowerOf2(n int) bool {
	return n&(n-1) == 0
}
