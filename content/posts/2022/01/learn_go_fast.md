---
author: "jdlau"
date: 2022-01-25
linktitle: Go快速入门
menu:
next:
prev:
title: Go快速入门
weight: 10
categories: ['Go']
tags: ['learn']
---

## 源码

```go
// 所有代码都需要放到包里
package color

// 导入其它包
import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 枚举
type Color int

// 常量
const (
	Red   Color = 1 // 红
	Blue  Color = 2 // 蓝
	Green Color = 3 // 绿
)

// 函数
func NewCar(
	name string,
	rate int,
) *Car {
	return &Car{
		name: name,
		rate: rate,
	}
}

// 类型
type Car struct {
	// 类型字段
	name string // 首字母小写，非导出，只能包内使用
	rate int
}

// 类型方法
func (car *Car) String() string { // 首字母大写，导出，可供其它包使用
	return "[Car] name: " + car.name + ", rate: " + strconv.Itoa(car.rate) + "."
}

func (car *Car) Run(
	ctx context.Context, // 使用ctx实现超时控制
) {
	// 定时器，每隔rate秒执行一次
	ticker := time.NewTicker(time.Duration(car.rate) * time.Second)
	defer ticker.Stop() // defer语句，在方法退出前执行，做收尾工作

	// for range ticker.C { // 循环，遍历chan
	// 	fmt.Printf("%s\n", car)
	// }

	for {
		select {
		case <-ticker.C:
			{ // 代码块，让逻辑更聚合，更清晰
				timesMutex.Lock()
				count := 1
				if v, ok := times[car.name]; ok {
					count = v + 1
				}
				times[car.name] = count
				timesMutex.Unlock()
			}

			fmt.Printf("%s\n", car)

		case <-ctx.Done():
			return
		}
	}
}

// 接口
type Runner interface {
	Run(ctx context.Context)
}

// 变量
var (
	// 确保*Car实现了Runner接口
	_ Runner = (*Car)(nil)

	timesMutex = new(sync.RWMutex)       // 读写锁，唯一写，多个读，读时无写
	times      = make(map[string]int, 2) // 记录Car Run的次数；在声明时初始化，并配置容量
)
```

## 测试

```go
package color

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCar(t *testing.T) {
	// 超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 并发执行
	wg := new(sync.WaitGroup)

	for _, car := range []Runner{ // 遍历切片
		NewCar("lanbo", 2),
		NewCar("boshi", 3),
	} {
		wg.Add(1) // 记录一个
		go func(car Runner) {
			defer wg.Done() // 完成一个

			t.Run(car.(*Car).name, func(t *testing.T) { // 对接口断言，获得具体类型
				car.Run(ctx)
			})
		}(car)
	}

	// 等上面均完成
	wg.Wait()

	timesMutex.RLock()
	fmt.Printf("times: %+v\n", times)
	timesMutex.RUnlock()
}
```

## 执行

编译：`go build`

执行：`go test -v`；可以尝试添加`-race`标志检测代码是否存在数据竞争
