package chanx

import (
	"testing"
)

func Test_oneSendMultiRecv(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2 will recv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneSendMultiRecv()
		})
	}
}

func Test_multiSendMultiRecv(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2 print before 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiSendMultiRecv()
		})
	}
}
