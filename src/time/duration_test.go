package test

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	// // 要计算持续时间中的单位数量，请划分：
	second := time.Second
	t.Log(int64(second / time.Millisecond)) // prints 1000

	// 要将整数个单位转换为持续时间，请乘以：
	seconds := 10
	t.Log(time.Duration(seconds) * time.Second) // prints 10s

	// 持续时间
	t0 := time.Now()
	time.Sleep(3 * time.Second)
	t1 := time.Now()
	t.Log("持续时间：", t1.Sub(t0))
}
