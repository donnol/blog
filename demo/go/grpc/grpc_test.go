package grpc

import "testing"

func TestGrpc(t *testing.T) {
	go Server()

	Client()
}
