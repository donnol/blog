package sched

import (
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "demo1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Demo()

			// 不添加睡眠的话，主线程执行完大概率就结束了
			time.Sleep(1 * time.Second)
		})
	}
}
