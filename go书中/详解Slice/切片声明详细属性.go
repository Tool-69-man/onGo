package main

import "fmt"

func main() {
	//var  slice  = make([]int,10,5)//length and capacity swapped  长度和容量需要交换
	var slice01 = make([]int, 5, 10) // 容量大于长度
	fmt.Println("声明", slice01)
	fmt.Println(len(slice01))
	fmt.Println(cap(slice01))
}
