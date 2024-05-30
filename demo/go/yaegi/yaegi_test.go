package yaegi

import (
	"fmt"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/donnol/do"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/traefik/yaegi/stdlib/unsafe"
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

func TestGin(t *testing.T) {
	src := `
	package main

	import (
		"fmt"
		"github.com/gin-gonic/gin"
	)
	
	func main() {
		fmt.Println("Hello from Yaegi!")
	
		router := gin.Default()
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		router.Run(":8080")
	}	
	`

	t.Run("未指定gopath", func(t *testing.T) {
		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols) // 引入标准库

		_, err := i.Eval(src)

		// 因为未引入gin库，所以会报错：
		want := `6:3: import "github.com/gin-gonic/gin" error: unable to find source related to: "github.com/gin-gonic/gin". Either the GOPATH environment variable, or the Interpreter.Options.GoPath needs to be set`
		do.Assert(t, err.Error(), want, diff.LineDiff(err.Error(), want))
	})

	t.Run("指定gopath但库不存在", func(t *testing.T) {
		i := interp.New(interp.Options{GoPath: "/home/jd/go1"})
		i.Use(stdlib.Symbols) // 引入标准库

		_, err := i.Eval(src)
		want := `6:3: import "github.com/gin-gonic/gin" error: unable to find source related to: "github.com/gin-gonic/gin"`
		do.Assert(t, err.Error(), want, diff.LineDiff(err.Error(), want))
	})

	t.Run("生成库符号并使用", func(t *testing.T) {
		i := interp.New(interp.Options{
			// GoPath: "/home/jd/go",
		})
		i.Use(stdlib.Symbols) // 引入标准库
		i.Use(unsafe.Symbols) // 支持unsafe
		i.Use(Symbols)        // 先在本目录go get github.com/gin-gonic/gin, 然后yaegi extract github.com/gin-gonic/gin生成符号，再Use

		_, err := i.Eval(src)
		if err != nil {
			t.Error(err)
		}
	})
}
