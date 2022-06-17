package new_make

import (
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	type TC[T any] struct {
		name string
		want T
	}
	tests := []TC[A]{
		{name: "", want: A{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make[A](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() = %v, want %v", got, tt.want)
			}
		})
	}

	// make(chan int) is not nil
	// Make[chan int]() is nil
	// 	chantests := []TC[chan int]{
	// 		{name: "", want: make(chan int)},
	// 	}
	// 	for _, tt := range chantests {
	// 		t.Run(tt.name, func(t *testing.T) {
	// 			if got := Make[chan int](); !reflect.DeepEqual(got, tt.want) {
	// 				t.Errorf("Make() = %v, want %v", got, tt.want)
	// 			}
	// 		})
	// 	}
}
