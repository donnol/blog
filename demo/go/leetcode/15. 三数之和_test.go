package leetcode

import (
	"fmt"
	"testing"

	"github.com/donnol/do"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2}, {-1, 0, 1},
			},
		},
		{
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.args.nums), func(t *testing.T) {
			got := threeSum(tt.args.nums)
			for i, want := range tt.want {
				do.AssertSlice(t, got[i], want)
			}
		})
	}
}
