package str

import (
	"strings"
	"testing"

	"github.com/donnol/do"
)

func TestNext(t *testing.T) {
	ns := next("ddffeeffdd")
	do.AssertSlice(t, ns, []int{0, 1, 0, 0, 0, 0, 0, 0, 1, 2})

	// 一样的
	{
		ns := next("dddd")
		do.AssertSlice(t, ns, []int{0, 1, 2, 3})
	}
	// 不存在一样的
	{
		ns := next("dfeg")
		do.AssertSlice(t, ns, []int{0, 0, 0, 0})
	}
	// 存在一样，但没有公共前后缀的
	{
		ns := next("dffg")
		do.AssertSlice(t, ns, []int{0, 0, 0, 0})
	}
	// 存在一样，有公共前后缀的
	{
		ns := next("dffd")
		// df -> d; f
		// dff -> d, df; ff, f
		// dffd -> d, df, dff; ffd, fd, d
		do.AssertSlice(t, ns, []int{0, 0, 0, 1})
	}
	{
		ns := next("ababac")
		do.AssertSlice(t, ns, []int{0, 0, 1, 2, 3, 0})
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		s       string
		pattern string
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
				s:       "abcdfdfdfdf",
				pattern: "dfd", // pattern dfd 's next array [0 0 1]
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				s:       "abcdefghijk",
				pattern: "ijk", // pattern ijk 's next array [0 0 0]
			},
			want: 8,
		},
		{
			name: "3",
			args: args{
				s:       "abcdefghijkOKM1232323232dfdfdfEFEFEFAERER",
				pattern: "AER", // pattern AER 's next array [0 0 0]
			},
			want: 36,
		},
		{
			name: "4",
			args: args{
				s:       "abcdefghijkOKM1232323232dfdfdfEFEFEFAERER",
				pattern: "dcb", // pattern dcb 's next array [0 0 0]
			},
			want: -1,
		},
		{
			name: "5",
			args: args{
				s:       "abcdefghijkOKM1232323232dfdfdfEFEFEFAERER",
				pattern: "abc", // pattern abc 's next array [0 0 0]
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{
				var got int
				if got = Index(tt.args.s, tt.args.pattern); got != tt.want {
					t.Errorf("Index() = %v, want %v", got, tt.want)
				}
				if got != -1 {
					if tt.args.s[got:got+len(tt.args.pattern)] != tt.args.pattern {
						t.Errorf("string slice = %v, want %v, ", tt.args.s[got:got+len(tt.args.pattern)], tt.args.pattern)
					}
				}
			}
			{
				var got int
				if got = strings.Index(tt.args.s, tt.args.pattern); got != tt.want {
					t.Errorf("Index() = %v, want %v", got, tt.want)
				}
				if got != -1 {
					if tt.args.s[got:got+len(tt.args.pattern)] != tt.args.pattern {
						t.Errorf("string slice = %v, want %v, ", tt.args.s[got:got+len(tt.args.pattern)], tt.args.pattern)
					}
				}
			}
		})
	}
}

// BenchmarkIndex-4
// 15456728        76.65 ns/op      24 B/op       1 allocs/op
func BenchmarkIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Index(
			"abcdefghijkOKM1232323232dfdfdfEFEFEFAERER",
			"AER",
		)
	}
}

// BenchmarkStdIndex-4
// 155435454         7.732 ns/op       0 B/op       0 allocs/op
func BenchmarkStdIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Index(
			"abcdefghijkOKM1232323232dfdfdfEFEFEFAERER",
			"AER",
		)
	}
}
