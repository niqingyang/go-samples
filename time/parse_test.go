package test

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	// https://www.jianshu.com/p/c7f7fbb16932
	// https://segmentfault.com/q/1010000010976398/a-1020000010982052
	// 月 1, 01, Jan, January
	// 日　 2, 02, _2
	// 时　 3, 03, 15, PM, pm, AM, am
	// 分　 4, 04
	// 秒　 5, 05
	// 年　 06, 2006
	// 时区 -07,-0700,Z0700,Z07:00,-07:00,MST
	// 周 Mon,Monday

	t0, _ := time.Parse("2006-01-02 15:04:05 -07", "2019-08-14 15:12:30 +08")

	t.Log(t0.Format("2006-01-02 15:04:05 -07"))
}
