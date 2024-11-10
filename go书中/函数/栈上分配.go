package main

import "fmt"

func main() {
	var b *int = test()
	fmt.Println(b, &b)
}

// 从函数返回 局部变量指针是安全的
func test() *int {
	a := 0x100
	return &a
}
