package main

import "fmt"

const (
	o = 1
	k = 2
	m = 3
)
const i, j, n = 12, 34, 56
const i2, j2, n2 = 12, 34, 56 //常量可以不引用

const (
	x = 1
	y
	l
)

func main() {
	fmt.Println(o, " ", k, " ", m)
	fmt.Println(i, " ", j, " ", n)
	fmt.Println(x, " ", y, " ", l) // 全是1

}
