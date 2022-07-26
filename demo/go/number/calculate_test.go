package number

import "testing"

func TestMul1(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "", args: args{a: 12.11, b: 3.24, c: 0.92}, want: 36.10},
		{name: "", args: args{a: 13.22, b: 3.24, c: 0.92}, want: 39.41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul1(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Mul1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul2(t *testing.T) {
	type args struct {
		b  float64
		c  float64
		as [3]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "", args: args{as: [3]float64{0.1, 1.1, 10.91}, b: 3.24, c: 0.92}, want: 36.10},
		{name: "", args: args{as: [3]float64{0.2, 1.3, 11.72}, b: 3.24, c: 0.92}, want: 39.41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul2(tt.args.b, tt.args.c, tt.args.as); got != tt.want {
				t.Errorf("Mul1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul3(t *testing.T) {
	type args struct {
		b  float64
		c  float64
		as [3]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "", args: args{as: [3]float64{0.1, 1.1, 10.91}, b: 3.24, c: 0.92}, want: 36.10},
		{name: "", args: args{as: [3]float64{0.2, 1.3, 11.72}, b: 3.24, c: 0.92}, want: 39.41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul3(tt.args.b, tt.args.c, tt.args.as); got != tt.want {
				t.Errorf("Mul1() = %v, want %v", got, tt.want)
			}
		})
	}
}
