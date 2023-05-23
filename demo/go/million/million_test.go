package million

import "testing"

func Test_start(t *testing.T) {
	type args struct {
		numRoutines int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1k",
			args: args{
				numRoutines: 1_000,
			},
		},
		{
			name: "10k",
			args: args{
				numRoutines: 10_000,
			},
		},
		{
			name: "100k",
			args: args{
				numRoutines: 100_000,
			},
		},
		{
			name: "1m",
			args: args{
				numRoutines: 1_000_000,
			},
		},
	}
	t.Log(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start(tt.args.numRoutines)
		})
	}
}
