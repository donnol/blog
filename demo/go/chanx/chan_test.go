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
			// 2 i: 1
		})
	}
}

func Test_oneSendManyMultiRecv(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2 print before 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneSendManyMultiRecv()
			// 2 i: 1
			// 1 i: 2
		})
	}
}

func Test_multiSendMultiRecv(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiSendMultiRecv()
			// 1 i: 2
			// 2 i: 1
		})
	}
}
