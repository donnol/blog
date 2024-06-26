package leetcode

import (
	"reflect"
	"testing"

	"github.com/donnol/do"
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

func TestSolution3(t *testing.T) {
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
			if got := Solution3(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution3(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^BenchmarkSolution$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// BenchmarkSolution3
	// BenchmarkSolution3-12    	24708423	        47.01 ns/op	      80 B/op	       1 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.216s

	for i := 0; i < b.N; i++ {
		Solution3(10)
	}
}

func TestSolution4(t *testing.T) {
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
			if got := Solution4(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution4(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^BenchmarkSolution$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// BenchmarkSolution4
	// BenchmarkSolution4-12    	21492500	        51.69 ns/op	      80 B/op	       1 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.176s

	for i := 0; i < b.N; i++ {
		Solution4(10)
	}
}

func Test_amountOf1V4(t *testing.T) {
	for _, tt := range []struct {
		n    uint32
		want int
	}{
		{
			n:    10000,
			want: 5,
		},
		{
			n:    10000000,
			want: 8,
		},
		{
			n:    1000000000,
			want: 13,
		},
	} {
		r := amountOf1V4(tt.n)
		do.Assert(t, r, tt.want)
	}
}

func Benchmark_amountOf1V4(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^Benchmark_amountOf1V4$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// Benchmark_amountOf1V4
	// Benchmark_amountOf1V4-12    	189971991	         6.462 ns/op	       0 B/op	       0 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.872s

	for i := 0; i < b.N; i++ {
		amountOf1V4(uint32(i))
	}
}

func TestSolution5(t *testing.T) {
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
			if got := Solution5(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution5(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^BenchmarkSolution$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// BenchmarkSolution5
	// BenchmarkSolution5-12    	29110830	        41.06 ns/op	      80 B/op	       1 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.245s

	for i := 0; i < b.N; i++ {
		Solution5(10)
	}
}

func TestSolution6(t *testing.T) {
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
			if got := Solution6(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution6(b *testing.B) {
	// go test -benchmem -run=^$ -tags linux -bench ^BenchmarkSolution$ github.com/donnol/blog/demo/go/leetcode -v

	// goos: linux
	// goarch: amd64
	// pkg: github.com/donnol/blog/demo/go/leetcode
	// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	// BenchmarkSolution6
	// BenchmarkSolution6-12    	28875668	        39.14 ns/op	      80 B/op	       1 allocs/op
	// PASS
	// ok  	github.com/donnol/blog/demo/go/leetcode	1.179s

	for i := 0; i < b.N; i++ {
		Solution6(10)
	}
}

func TestShift(t *testing.T) {

	a := 0b1
	t.Logf("%d, %06b; %b", a, a, a<<62)
	t.Log(1<<2, 1<<3, 2&(2-1) == 0, 3&(3-1) == 0, IsPowerOf2(2), IsPowerOf2(3))

	r := amountOf1V3(5)
	do.Assert(t, r, 2)
}
