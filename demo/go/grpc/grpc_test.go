package grpc

import (
	"testing"
	"time"
)

func TestGrpc(t *testing.T) {
	go Server()

	time.Sleep(100 * time.Millisecond)
	Client()
}
