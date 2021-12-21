/*
 * Copyright (c) 2017 - love Co.,Ltd. All rights reserved.
 *
 * @Author: jdlau
 * @Date: 2021-12-08 17:05:14
 * @Last Modified by: jdlau
 * @Last Modified time: 2021-12-08 17:46:55
 */
package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func Server() {
	server := grpc.NewServer()

	defer server.GracefulStop()

	info := server.GetServiceInfo()
	fmt.Printf("info: %+v\n", info)

	// 可以调多次，注册多个服务
	server.RegisterService(&grpc.ServiceDesc{
		ServiceName: "pinger",
		HandlerType: (*Pinger)(nil), // 接口指针
		Methods: []grpc.MethodDesc{
			{
				MethodName: "ping",
				Handler: func(srv interface{}, ctx context.Context, f func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					// srv是绑定的具体实例，如果有多个方法时，怎么对应呢？
					// Methods是一个数组，多个方法对应到数组的元素里
					fmt.Printf("srv: %v\n", srv)

					fmt.Printf("ctx: %+v\n", ctx)

					// 参数在哪里呢？
					// 参数需要通过调用f来获取，如下：
					// in := new(Param)
					// if err := dec(in); err != nil {
					// 	return nil, err
					// }

					// 这里执行的f又是哪里来的呢？执行它又有什么用呢？
					// 如上所述，f是用来获取参数的
					var m = message("hello")
					if err := f(&m); err != nil {
						panic(err)
					}

					r := message("good")
					return &r, nil
				},
			},
		},
		Streams:  nil,
		Metadata: nil,
	}, &pingImpl{name: "jd"})

	l, err := net.Listen("tcp", "127.0.0.1:8980")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}

type Pinger interface {
	Ping(msg string) string
}

type pingImpl struct {
	name string
}

func (p *pingImpl) Ping(msg string) string {
	fmt.Printf("msg: %s\n", msg)

	return "I am " + p.name
}

func Client() {
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8980", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	m := message("hello-from-client")
	var r message
	if err := conn.Invoke(context.Background(), "/pinger/ping", &m, &r); err != nil { // /pinger/ping与上述注册的服务方法对应
		panic(err)
	}

	fmt.Printf("r: %s\n", r)
}

type message string

func (m *message) Reset() {
}

func (m *message) String() string {
	return string(*m)
}

func (m *message) ProtoMessage() {

}

func (m *message) Marshal() ([]byte, error) {
	return []byte(*m), nil
}

func (m *message) Unmarshal(data []byte) error {
	*m = message(data)
	return nil
}
