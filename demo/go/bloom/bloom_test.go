package bloom

import "testing"

func TestBloom(t *testing.T) {
	type args struct {
		m    uint
		n    uint
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "", args: args{m: 10, n: 4, data: []byte(`123`)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Bloom(tt.args.m, tt.args.n, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Bloom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
