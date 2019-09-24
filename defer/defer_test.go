package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
