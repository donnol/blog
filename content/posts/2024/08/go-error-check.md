```go
func Call(a, b, c, d, e any) (any, error) {
    // 1
    v, err := call(a)
    check err // if err != nil { return }

    // Nomarlly, `check` will return if err != nil, it will continue if err == nil.
    // the `return` will auto match the function's result
    // 
    // If we want to handle err, we can use a code block to do that.

    // 2
    v, err := call(b)
    check err { log.Printf("call failed: %v", err) } // print error and continue, because no `return` in block

    // 3
    v, err := call(c)
    check err { log.Printf("call failed: %v", err); return } // print error and return, because `return` in block

    // 4
    v, err := call(d)
    check err { err = fmt.Errorf("call failed: %w", err); return } // wrap error and return

    // 5
    v, err := call(e)
    check err { panic(err) } // panic
}
```

整下来，跟`if err != nil { ... }`也差不多。

```go
func Call() {
    v, err := call(a) 

    v := Log(call(a))
    v := Must(call(a))
    v := Wrap(call(a), "call failed: %w")

    v := try(call(a), "call failed: %w") // 这么一看，内置`try`函数不也挺好

    v := try(call(a), handler) // 仅当call(a)返回的err不为nil时执行err handler
    // err handler可以Log、Wrap、Panic等
}
```

`try`是内置函数，在编译时展开；支持任意数量的参数和任意数量的返回值，只对最后一个返回是error进行处理，处理时根据设置的handler执行，未设置handler时默认return。
