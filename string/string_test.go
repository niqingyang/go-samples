package test

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestLen(t *testing.T) {
	s := "hello,世界"
	r := []rune(s)

	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%p\n", &s)

	t.Log(s[7:])
	t.Log(utf8.DecodeRuneInString(s[7:]))
	t.Log(len(s))
	t.Log(utf8.RuneCountInString(s))



	for i, rs := range r {
		t.Logf("%d %q %v", i, rs, rs)
	}

}

// golang 中 string 底层是通过 byte 数组实现的
// 中文字符在 unicode 下占 2 个字节，在 utf-8 编码下占 3 个字节，而 golang 默认编码正好是 utf-8
func TestStrLen(t *testing.T)  {
	s := "hello,世界" // byte 是 uint8 的别名
	r := []rune(s) // runte 是 int32 的别名，所以此处 r := []int32(s) 效果一样

	// golang 中 string 底层是通过 byte 数组实现的，座椅直接求 len 实际是在按字节长度计算，所以一个汉字占3个字节算了3个长度
	t.Logf("按 []byte 数组进行计算长度：%v", len(s))
	t.Logf("按 []rune 数组进行计算长度：%v", len(r))
	t.Logf("通过 utf8 函数获取字符串长度：%v", utf8.RuneCountInString(s))
	t.Logf("分割字符串：%v", strings.Split(s, ""))
}

func TestStr(t *testing.T)  {
	s  := []rune("hello, 世界")

	t.Logf("%s", string(s[7:]))
}
