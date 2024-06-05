package str

import (
	"testing"
)

func TestToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{
				s: "123",
			},
			want: 123,
		},
		{
			name: "negative",
			args: args{
				s: "-123",
			},
			want: -123,
		},
		{
			name: "full",
			args: args{
				s: "1234567890",
			},
			want: 1234567890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.s); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt_Panic(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "input is empty",
		},
		{
			name: "invalid",
			args: args{
				s: "-123c",
			},
			want: "invalid char c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(want string) {
				if v := recover(); v != nil {
					e := v.(error).Error()
					if e != want {
						t.Errorf("bad case: %v != %v", e, want)
					}
				}
			}(tt.want)
			ToInt(tt.args.s)
		})
	}
}
