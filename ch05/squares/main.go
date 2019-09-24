// squares 函数返回一个函数，后者包含下一次要用到的平方数
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++ // 匿名函数能够获取到整个词法环境，因此里层的函数可以使用外层函数中的变量
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
}
