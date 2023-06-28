package rotate

import (
	"reflect"
	"testing"
)

func TestFront(t *testing.T) {
	type args struct {
		s     []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "s",
			args: args{
				s:     []int{1, 2, 3},
				index: 0,
			},
			want: []int{2, 3, 0},
		},
		{
			name: "m",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 2,
			},
			want: []int{1, 2, 4, 6, 7, 0},
		},
		{
			name: "eb",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 5,
			},
			want: []int{1, 2, 3, 4, 6, 0},
		},
		{
			name: "e",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 6,
			},
			want: []int{1, 2, 3, 4, 6, 7},
		},
		{
			name: "eg",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 7,
			},
			want: []int{1, 2, 3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Front(tt.args.s, tt.args.index)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("bad case: %v != %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		s     []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "s",
			args: args{
				s:     []int{1, 2, 3},
				index: 0,
			},
			want: []int{2, 3},
		},
		{
			name: "m",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 2,
			},
			want: []int{1, 2, 4, 6, 7},
		},
		{
			name: "eb",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 5,
			},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "e",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 6,
			},
			want: []int{1, 2, 3, 4, 6, 7},
		},
		{
			name: "eg",
			args: args{
				s:     []int{1, 2, 3, 4, 6, 7},
				index: 7,
			},
			want: []int{1, 2, 3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Remove(&tt.args.s, tt.args.index)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("bad case: %v != %v", tt.args.s, tt.want)
			}
		})
	}
}
