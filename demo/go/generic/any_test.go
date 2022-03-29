package generic

import (
	"reflect"
	"testing"
)

func TestUseAnyDirect(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"different", args{a: 1, b: "2"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseAnyDirect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseAnyDirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseAnyGeneric(t *testing.T) {
	type args[T, R any] struct {
		a T
		b R
	}
	type testCase[T, R any] struct {
		name string
		args args[T, R]
		want R
	}
	tests := []testCase[int, string]{
		{"", args[int, string]{1, "1"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseAnyGeneric(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseAnyGeneric() = %v, want %v", got, tt.want)
			}
		})
	}
}
