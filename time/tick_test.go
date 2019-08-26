// Tick 是 NewTicker 的便捷包装，仅提供对滴答声频道的访问。
// 虽然 Tick 对于不需要关闭 Ticker 的客户端非常有用，但请注意，
// 如果没有办法关闭它，底层 Ticker 不能被垃圾收集器恢复;
// 它“泄漏”。与 NewTicker 不同，如果 d <= 0，Tick 将返回 nil。
package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v\n", now)
	}
}
