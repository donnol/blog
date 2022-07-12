package stack

import (
	"math/rand"
	"testing"

	"github.com/donnol/blog/demo/go/fmtx"
)

func TestNormalStack(t *testing.T) {
	fmtx.Enable = true

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "10", args: args{n: 10}},
	}
	for _, tt := range tests {
		arr := make([]int, 0, tt.args.n)
		for i := 0; i < tt.args.n; i++ {
			arr = append(arr, rand.Intn(100))
		}
		t.Run(tt.name, func(t *testing.T) {
			NormalStack(arr)
		})
	}
}

func TestMonoIncrStack(t *testing.T) {
	fmtx.Enable = true

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "10", args: args{n: 10}},
	}
	for _, tt := range tests {
		arr := make([]int, 0, tt.args.n)
		for i := 0; i < tt.args.n; i++ {
			arr = append(arr, rand.Intn(100))
		}
		t.Run(tt.name, func(t *testing.T) {
			MonoIncrStack(arr)
		})
	}
}

var (
	arr = func() []int {
		n := 10000
		arr := make([]int, 0, n)
		for i := 0; i < n; i++ {
			arr = append(arr, rand.Intn(100))
		}
		return arr
	}()
)

func BenchmarkMonoIncrStack(b *testing.B) {
	fmtx.Enable = false

	for i := 0; i < b.N; i++ {
		MonoIncrStack(arr)
	}
}
