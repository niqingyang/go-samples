// 在等待持续时间后，然后在返回的频道上发送当前时间。它相当于NewTimer(d).C。
// 在定时器触发之前，底层 Timer 不会被垃圾收集器恢复。如果需要提高效率，请改用 NewTimer ，
// 如果不再需要定时器，则调用 Timer.Stop。
package test

import (
	"testing"
	"time"
)

func Test(t *testing.T)  {
	t.Log(<-time.After(time.Second * 3))
}