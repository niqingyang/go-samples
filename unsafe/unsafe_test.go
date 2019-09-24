package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeOf(t *testing.T) {
	fmt.Println(unsafe.Sizeof(float64(0)))
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func TestFloat64bits(t *testing.T) {
	fmt.Printf("%#016x\n", Float64bits(1.0))
}
