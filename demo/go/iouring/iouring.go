//go:build linux

package iouring

import (
	"fmt"
	"os"

	"github.com/iceber/iouring-go"
)

func IOURing() {
	// go get github.com/iceber/iouring-go
	var str = "io with iouring"

	iour, err := iouring.New(1)
	if err != nil {
		panic(fmt.Sprintf("new IOURing error: %v", err))
	}
	defer iour.Close()

	file, err := os.Create("./tmp.txt")
	if err != nil {
		panic(err)
	}

	// iouring与文件fd关联，请求做一个写str到文件的操作
	prepRequest := iouring.Write(int(file.Fd()), []byte(str))

	// 提交请求，传入req函数和ch管道
	ch := make(chan iouring.Result, 1)
	if _, err := iour.SubmitRequest(prepRequest, ch); err != nil {
		panic(err)
	}

	// 通过ch管道，等待结果返回
	result := <-ch

	// 解析结果
	fd, err := result.ReturnFd()
	if err != nil {
		fmt.Println("return fd failed: ", err)
		return
	}
	fmt.Printf("return fd: %v\n", fd)

	n, err := result.ReturnInt()
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}
	fmt.Printf("write byte: %d\n", n)
}
