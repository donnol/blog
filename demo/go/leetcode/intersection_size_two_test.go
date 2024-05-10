package leetcode

import (
	"reflect"
	"testing"
)

func Test_intersectionSizeTwo(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionSizeTwo(tt.args.intervals); got != tt.want {
				t.Errorf("intersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect_2(t *testing.T) {
	type args struct {
		curr []int
		next []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{
				curr: []int{1, 2},
				next: []int{3, 7},
			},
			want: []int{},
		},
		{
			name: "1",
			args: args{
				curr: []int{1, 3},
				next: []int{3, 7},
			},
			want: []int{3},
		},
		{
			name: "2",
			args: args{
				curr: []int{3, 4},
				next: []int{3, 7},
			},
			want: []int{3, 4},
		},
		{
			name: "3",
			args: args{
				curr: []int{1, 2},
				next: []int{3, 4},
			},
			want: []int{},
		},
		{
			name: "4",
			args: args{
				curr: []int{3, 4},
				next: []int{1, 2},
			},
			want: []int{},
		},
		{
			name: "4",
			args: args{
				curr: []int{3, 8},
				next: []int{1, 8},
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.curr, tt.args.next, 2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
