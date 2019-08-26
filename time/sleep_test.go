package test

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	t.Log("--- 开始等待")
	time.Sleep(3 * time.Second)
	t.Log("--- 等待结束")
}
