// 输出命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 是否忽略正常输出结尾的换行符
var n = flag.Bool("n", false, "omit trailing newline")
// 使用 sep 替换默认参数输出时使用的空格分隔符
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
