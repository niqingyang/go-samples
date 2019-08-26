package test

import (
	"fmt"
	"testing"
	"unicode"
)

func TestUnicode(t *testing.T) {

	fmt.Print("\a")

	// 判断字符是否为十进制的整数
	t.Log(unicode.IsDigit(rune('1')), unicode.IsDigit(rune('A')), unicode.IsDigit(rune('a')))
	// 判断字符是否为字母
	t.Log(unicode.IsLetter(rune('1')), unicode.IsLetter(rune('A')), unicode.IsLetter(rune('a')))
}
