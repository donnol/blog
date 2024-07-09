package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abbd",
			args: args{
				s: "abbd",
			},
			want: "bb",
		},
		{
			name: "abba",
			args: args{
				s: "abba",
			},
			want: "abba",
		},
		{
			name: "abab",
			args: args{
				s: "abab",
			},
			want: "aba",
		},
		{
			name: "ccc",
			args: args{
				s: "ccc",
			},
			want: "ccc",
		},
		{
			name: "aaaa",
			args: args{
				s: "aaaa",
			},
			want: "aaaa",
		},
		{
			name: "aaaaa",
			args: args{
				s: "aaaaa",
			},
			want: "aaaaa",
		},
		{
			name: "abb",
			args: args{
				s: "abb",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
