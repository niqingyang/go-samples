package _interface

import (
	"bytes"
	"io"
	"testing"
)

func TestNilInterface(t *testing.T) {
	var buf *bytes.Buffer
	var buf1 *bytes.Buffer

	var w = io.Writer(buf)
	var w1 = io.Writer(buf1)

	t.Logf("%v", w)
	t.Log(w == nil)
	t.Log(w == w1)

	var i io.Writer

	t.Logf("%v", i)
	t.Log(i == nil)
}
