package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				n: 10,
			},
			want: []int{
				0, 1, 1, 2, 1, 2, 2, 3, 1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^BenchmarkSolution$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// BenchmarkSolution
	// BenchmarkSolution-12    	 1735036	       710.3 ns/op	     106 B/op	       9 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.937s

	for i := 0; i < b.N; i++ {
		Solution(10)
	}
}

func TestSolution2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				n: 10,
			},
			want: []int{
				0, 1, 1, 2, 1, 2, 2, 3, 1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution2(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^BenchmarkSolution$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// BenchmarkSolution
	// BenchmarkSolution2-12    	 5050453	       239.2 ns/op	     106 B/op	       9 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.454s

	for i := 0; i < b.N; i++ {
		Solution2(10)
	}
}
