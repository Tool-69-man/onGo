package main

import "fmt"

func main() {
	var a string = `\n1\不作转义处理
									并且支持跨行`
	fmt.Println(a)

	b := "ab" + //跨行必须 '+' 在上一行结尾
		"cd"
	fmt.Println(b == "abcd")
	fmt.Println(b > "abc")
}
