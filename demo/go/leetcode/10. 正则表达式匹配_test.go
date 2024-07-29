package leetcode

import (
	"fmt"
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "aa",
			args: args{
				s: "aa",
				p: "aa",
			},
			want: true,
		},
		{
			name: "aa",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "aa",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "ab",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "aab",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "ab",
			args: args{
				s: "ab",
				p: ".*c",
			},
			want: false,
		},
		{
			name: "aaa",
			args: args{
				s: "aaa",
				p: "aaaa",
			},
			want: false,
		},
		{
			name: "aaca",
			args: args{
				s: "aaca",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "a",
			args: args{
				s: "a",
				p: "ab*",
			},
			want: true,
		},
		{
			name: "a",
			args: args{
				s: "a",
				p: "ab*a",
			},
			want: false,
		},
		{
			name: "bbbba",
			args: args{
				s: "bbbba",
				p: ".*a*a",
			},
			want: true,
		},
		{
			name: "aaa",
			args: args{
				s: "aaa",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "a",
			args: args{
				s: "a",
				p: ".*..a*",
			},
			want: false,
		},
		{
			name: "ab",
			args: args{
				s: "ab",
				p: ".*..",
			},
			want: true,
		},
		{
			name: "a",
			args: args{
				s: "a",
				p: "a.",
			},
			want: false,
		},
		{
			name: "a",
			args: args{
				s: "a",
				p: ".*.",
			},
			want: true,
		},
		{
			name: "abbbcd",
			args: args{
				s: "abbbcd",
				p: "ab*bbbcd",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("isMatch(%s,%s)", tt.args.s, tt.args.p), func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch(%s,%s) = %v, want %v", tt.args.s, tt.args.p, got, tt.want)
			}
		})
	}
}
