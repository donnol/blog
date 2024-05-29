package leetcode

import (
	"testing"
)

func TestMaxLen(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				s: []string{"abc", "ab", "c", "cde", "efg"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxLen(tt.args.s); got != tt.want {
				t.Errorf("MaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProduct(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				s: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct(tt.args.s); got != tt.want {
				t.Errorf("MaxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
