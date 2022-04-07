package chanx

import (
	"testing"
	"time"
)

func Test_bufChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bufChan()

			time.Sleep(10 * time.Second)
		})
	}
}
