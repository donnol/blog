package leetcode

import (
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: "I",
		},
		{
			name: "2",
			args: args{
				num: 2,
			},
			want: "II",
		},
		{
			name: "5",
			args: args{
				num: 5,
			},
			want: "V",
		},
		{
			name: "10",
			args: args{
				num: 10,
			},
			want: "X",
		},
		{
			name: "4",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "9",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "3749",
			args: args{
				num: 3749,
			},
			want: "MMMDCCXLIX",
		},
		{
			name: "58",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "1994",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binPow(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1^2",
			args: args{
				a: 1,
				b: 2,
			},
			want: 1,
		},
		{
			name: "3^4",
			args: args{
				a: 3,
				b: 4,
			},
			want: 81,
		},
		{
			name: "10^0",
			args: args{
				a: 10,
				b: 0,
			},
			want: 1,
		},
		{
			name: "10^2",
			args: args{
				a: 10,
				b: 2,
			},
			want: 100,
		},
		{
			name: "10^8",
			args: args{
				a: 10,
				b: 8,
			},
			want: 100000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("binPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
