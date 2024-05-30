// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type I int
type R string

type I2 int
type R2 string

func Do[P, R any](p P) (r R) {
	if true {
		fmt.Printf("%T, %v\n", r, reflect.TypeOf(r))
	}
	return
}

func main() {
	fmt.Println("Hello, 世界")

	Do[I, R](1)
	Do[I2, R2](1)
	Do[I, R](1)
}
