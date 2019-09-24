// 利用二叉树实现插入排序
package main

import "fmt"

// 结构体
type tree struct {
	value       int
	left, right *tree
}

// 排序
func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

// 将元素按照顺序追加到 values 中，然后返回结果
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 插入值时定位值唯一节点的做或者右
func add(t *tree, value int) *tree {
	if t == nil {
		// 等价于 t = &tree{value: value}
		t = new(tree)
		t.value = value
		return t;
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{63, 2, 1, 8, 10, 4, 8}
	fmt.Printf("%v", Sort(values))
}
