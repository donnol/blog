package heart

import (
	"testing"
)

func TestDraw(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Draw()
		})
	}
}
