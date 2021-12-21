package pb

import (
	context "context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	grpc "google.golang.org/grpc"
)

type ping struct{}

func (p *ping) Ping(ctx context.Context, param *Param) (*Result, error) {
	fmt.Printf("p: %+v\n", param)
	return &Result{
		Name: "Result",
	}, nil
}

func TestNewPingerClient(t *testing.T) {
	// c := make(chan struct{})

	// start server
	go func() {
		server := grpc.NewServer()
		defer server.GracefulStop()

		RegisterPingerServer(server, &ping{})

		l, err := net.Listen("tcp", "127.0.0.1:8980")
		if err != nil {
			panic(err)
		}

		// Serve方法如果多接受一个chan参数，在可以Accept时往该chan写入值，从而通知外面它已经就绪，可以开始接受请求了，就不需要通过睡眠一小段时间来等待它就绪了
		// if err := server.ServeChan(l, c); err != nil {

		if err := server.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
	// <-c

	time.Sleep(1 * time.Second)

	// client
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8980", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := NewPingerClient(conn)
	r, err := client.Ping(context.Background(), &Param{
		Name: "Param",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("r: %+v\n", r)
}
