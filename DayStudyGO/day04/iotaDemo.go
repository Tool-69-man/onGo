package main

import "fmt"

const (
	a = iota //0
	b
	c
	d = "xiong"
	e
	f = 100
	g
	h = iota
	i
)

const (
	x1 = 1 << iota //0
	x2 = 2 << iota
	x3
	x4
)

func main() {
	fmt.Println(a, " 0")
	fmt.Println(b, " 1")
	fmt.Println(c, " 2")
	fmt.Println(d, " xiong")
	fmt.Println(e, " 4") //错误   还是xiong  由此可见如果没有iota,平时const就和上一个保持一致
	fmt.Println(f, " 100")
	fmt.Println(g, " 6") // 错误  还是100
	fmt.Println(h, " 7") //又将iota 赋值给常量
	fmt.Println(i, " 8")

	fmt.Println("up--------------")
	fmt.Println(x1) // 1<<0
	fmt.Println(x2) // 2 << 1
	fmt.Println(x3) // 2 << 2
	fmt.Println(x4) // 2 << 3

}
