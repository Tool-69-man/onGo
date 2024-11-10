package main

import "fmt"

// 每个引用类型都是interface{}
func main() {
	num := []int{1,2,3}
	num.
	var i = num.(int)
	fmt.Println(i)
}
