package math

import (
	"math/rand"
	"testing"
	"time"
)

func TestSeed(t *testing.T)  {
	// 只有设置 seed 后才会产生不同的随机数
	rand.Seed(time.Now().UTC().UnixNano());

	// Intn 以 int 形式返回来自默认 Source 的 [0，n] 中的非负伪随机数。如果 n <= 0，它会发生混乱。
	// rand.Intn(100) 返回 0 <= n <= 100
	t.Log(rand.Intn(100))
	t.Log(rand.Intn(100))
	t.Log(rand.Intn(100))


}

func TestNewSource(t *testing.T)  {
	// seed 值相同产生的随机数相同
	s1 := rand.NewSource(100)
	r1 := rand.New(s1)

	t.Log(r1.Intn(100))
	t.Log(r1.Intn(100))
	t.Log(r1.Intn(100))
}
