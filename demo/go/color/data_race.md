# data race

go test -race

```
[Car] name: lanbo, rate: 2.
==================
WARNING: DATA RACE
Read at 0x00c0000a6210 by goroutine 11:
  runtime.mapaccess2_faststr()
      /usr/local/go/src/runtime/map_faststr.go:107 +0x0
  github.com/donnol/blog/demo/go/color.(*Car).Run()
      /home/jd/Project/blog/demo/go/color/color.go:59 +0x136
  github.com/donnol/blog/demo/go/color.TestCar.func1.1()
      /home/jd/Project/blog/demo/go/color/color_test.go:28 +0x61
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Previous write at 0x00c0000a6210 by goroutine 10:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/donnol/blog/demo/go/color.(*Car).Run()
      /home/jd/Project/blog/demo/go/color/color.go:62 +0x1aa
  github.com/donnol/blog/demo/go/color.TestCar.func1.1()
      /home/jd/Project/blog/demo/go/color/color_test.go:28 +0x61
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Goroutine 11 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1306 +0x726
  github.com/donnol/blog/demo/go/color.TestCar.func1()
      /home/jd/Project/blog/demo/go/color/color_test.go:27 +0x1d9
  github.com/donnol/blog/demo/go/color.TestCar·dwrap·2()
      /home/jd/Project/blog/demo/go/color/color_test.go:30 +0x58

Goroutine 10 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1306 +0x726
  github.com/donnol/blog/demo/go/color.TestCar.func1()
      /home/jd/Project/blog/demo/go/color/color_test.go:27 +0x1d9
  github.com/donnol/blog/demo/go/color.TestCar·dwrap·2()
      /home/jd/Project/blog/demo/go/color/color_test.go:30 +0x58
==================
[Car] name: boshi, rate: 3.
[Car] name: lanbo, rate: 2.
[Car] name: boshi, rate: 3.
[Car] name: lanbo, rate: 2.
[Car] name: lanbo, rate: 2.
[Car] name: boshi, rate: 3.
times: map[boshi:3 lanbo:4]
--- FAIL: TestCar (10.00s)
    --- FAIL: TestCar/boshi (10.00s)
        testing.go:1152: race detected during execution of test
    --- FAIL: TestCar/lanbo (10.00s)
        testing.go:1152: race detected during execution of test
    testing.go:1152: race detected during execution of test
FAIL
exit status 1
FAIL    github.com/donnol/blog/demo/go/color    10.011s
```

多个线程都对`map`读写时，发生竞争；因为原生`map`的实现偏向通用，没有保证并发读写安全。

解决：

- 为`map`添加`sync.Mutex`或`sync.RWMutex`保护

- 使用`sync.Map`
