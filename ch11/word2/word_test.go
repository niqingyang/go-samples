package word

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`!IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	// 会输出多次，不断调整次数
	for i := 0; i < b.N; i++ {
		IsPalindrome("palindrome")
	}

}
