package noescape

import (
	"reflect"
	"testing"
	"unsafe"
)

func Test_noescape(t *testing.T) {
	type args struct {
		p unsafe.Pointer
	}
	p := new(int)
	tests := []struct {
		name string
		args args
		want unsafe.Pointer
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				p: unsafe.Pointer(p),
			},
			want: unsafe.Pointer(p),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noescape(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noescape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xor(t *testing.T) {
	for _, x := range []uint{
		1, 2, 3,
	} {
		r := x ^ 0 //nolint:staticcheck SA4016 x ^ 0 always equals x
		t.Logf("x ^ 0 = %d", r)

		if r != x {
			t.Errorf("bad case: %d != %d", r, x)
		}
	}
}
