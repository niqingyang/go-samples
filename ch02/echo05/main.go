package main

import "fmt"

func main() {
	x := "hello!"

	for i := 0; i < len(x); i++ {

		// 此处 x = "hello!"

		// 新的词法块，新的循环开始，x[i] 并不在此词法块中，所以使用外部词法块中的 x[i]
		// 赋值后，当前词法块中的 x 会覆盖掉外部词法块的 x，但不是将之前的 x 清除掉
		x := x[i]

		// 此处 x = 'h'

		if x != '!' {
			// 此处 x = 'h'
			// := 重新定义了变量 x，覆盖了外部词法块的变量
			x := x + 'A' - 'a'
			// 此处 x = 'H'
			fmt.Printf("%c\n", x)
		}

		// 此处 x = 'h'
		fmt.Printf("%c\n", x)
	}
}
