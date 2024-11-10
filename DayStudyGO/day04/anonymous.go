package main

import "fmt"

//匿名变量

func test() (int, int) {
	return 1, 2
}
func test2(a int) (int, int, bool) {
	return 1, 2, true
}

func main() {
	a, b := test()
	c, _, _ := test2(1)
	_, v, _ := test2(1)
	var _, _, g = test2(1)
	fmt.Println(a, "   ", b)
	fmt.Println(c)
	fmt.Println(v)
	fmt.Println(g)

}
