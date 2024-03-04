package yaegi

import (
	"fmt"
	"testing"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func TestYaegi(t *testing.T) {
	const src = `
package main

import (
	"fmt"
	"math/big"
	"time"
	"runtime"
)

func Sum(a, b int) int {
	return a + b
}

func SumG[T ~int](a, b T) T {
	return a + b
}

func Go() {
	go func() {
		for i:=0; i< 10;i++ {
			fmt.Println(i)
		}
	}()
	time.Sleep(1*time.Second)
}

func GoR() chan int {
	ch := make(chan int, 1)
	go func() {
		for i:=0; i< 10;i++ {
			ch<-i
		}
		close(ch)
	}()
	return ch
}

func call() {
	p, file, line, ok := runtime.Caller(0) // 拿到的位置信息不是脚本的
	fmt.Println("call by", p, file, line, ok)
}

func Call() {
	call()
}

// Factorial n <= 20, it will return 0 if n > 20, use FactorialBig instead.
func Factorial(n int) int {
	if n > 20 {
		return 0
	}

	s := 1
	for i := 2; i <= n; i++ {
		s *= i
	}
	return s
}

// FactorialBig n > 20
func FactorialBig(n int) string {
	s := big.NewInt(1)
	for i := 2; i <= n; i++ {
		s = s.Mul(s, big.NewInt(int64(i)))
	}
	return s.String()
}
`

	{
		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols) // 引入标准库

		_, err := i.Eval(src)
		if err != nil {
			t.Fatal(err)
		}

		// normal function
		v, err := i.Eval("main.Sum")
		if err != nil {
			t.Fatal(err)
		}

		bar := v.Interface().(func(int, int) int)
		r := bar(1, 1)
		fmt.Println(r)

		// generic function
		{
			v, err := i.Eval("main.SumG[int](1, 2)")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("v: %+v\n", v)
		}

		// goroutine
		{
			v, err := i.Eval("main.Go()")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("v: %+v\n", v.String())
		}

		// goroutine with result
		{
			v, err := i.Eval("main.GoR()")
			if err != nil {
				t.Fatal(err)
			}
			ch := v.Interface().(chan int)
			for v := range ch {
				fmt.Printf("v: %+v\n", v)
			}
		}

		// call
		{
			v, err := i.Eval("main.Call()")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("v: %+v\n", v.String())
		}

		// Factorial
		{
			v, err := i.Eval("main.Factorial(10)")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("v: %+v\n", v.Int())
		}
		{
			v, err := i.Eval("main.FactorialBig(30)")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("v: %+v\n", v.String())
		}
	}
}
