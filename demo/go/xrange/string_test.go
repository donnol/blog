package xrange

import "testing"

func TestRangeString(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   string
	}{
		{name: "same", in: "hello"},
		{name: "not same", in: "你好"},
	} {
		RangeString(tt.in)
		ForString(tt.in)
		// ascii字符串是相同的输出
		// 0, 104, h
		// 1, 101, e
		// 2, 108, l
		// 3, 108, l
		// 4, 111, o
		// 0, 104, h
		// 1, 101, e
		// 2, 108, l
		// 3, 108, l
		// 4, 111, o

		// 中文字符串时输出完全不一样
		// `range`是对每个字的遍历
		// 下标也不是连续的
		// 0, 20320, 你
		// 3, 22909, 好
		// `for`是对每个字节的遍历
		// 下标还是连续的
		// 0, 228, ä
		// 1, 189, ½
		// 2, 160,
		// 3, 229, å
		// 4, 165, ¥
		// 5, 189, ½
	}
}
