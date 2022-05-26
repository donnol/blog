package sqlserver

import (
	"testing"
)

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "demo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Demo()
		})
	}
}

func TestDemoGorm(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "demo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DemoGorm()
		})
	}
}
