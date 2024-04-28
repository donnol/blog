package rangestring

import "testing"

func TestRangeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				s: "abc",
			},
		},
		{
			name: "我是谁",
			args: args{
				s: "我是谁",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RangeString(tt.args.s)
		})
	}
}
