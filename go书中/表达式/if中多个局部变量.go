package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if a, b := x, y; a < b { //局部变量 有效范围包括整个if/else
		fmt.Println("a小")
	} else {
		fmt.Println("b小")
	}
	if a < 2 { //undefined: a
		fmt.Println("a已经无效")
	}
}
