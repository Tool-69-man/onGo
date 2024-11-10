package main

import (
	"flag"
	"fmt"
	"strings"
)

// 新的布尔标识是变量
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // 更新标记变量的的默认值
	fmt.Print("===", strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
