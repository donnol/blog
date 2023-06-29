package zero

import "testing"

func TestZero(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0",
			args: args{
				v: 0,
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				v: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero(tt.args.v); got != tt.want {
				t.Errorf("Zero() = %v, want %v", got, tt.want)
			}
		})
	}
}
