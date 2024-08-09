package leetcode

import (
	"fmt"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
		{
			args: args{
				nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
				target: -2,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.args.nums, tt.args.target), func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
