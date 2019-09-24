package hooks

import (
	"fmt"
	"testing"
)

type TestFilter struct {
}

func (f *TestFilter) Apply() interface{} {

	fmt.Println("done")

	return "test"
}

func TestMap(t *testing.T) {
	var hooks = new(Hooks)

	filter := new(TestFilter)

	hooks.AddFilter("test1", filter, 0)
	hooks.AddFilter("test2", filter, 0)

	hooks.ApplyFilter("test1")
	hooks.ApplyFilter("test2")
}
