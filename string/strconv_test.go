package test

import (
	"fmt"
	"strconv"
	"testing"
)

// 按不同的进制格式化数字
func TestFormatInt(t *testing.T) {
	x := 123
	y := fmt.Sprintf("%d", x)

	fmt.Println(y, strconv.Itoa(x))

	t.Log(fmt.Sprintf("x=%b", x), strconv.FormatInt(int64(x), 2))
}
