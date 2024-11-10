package main

import "fmt"

func main() {

	var sliceInt0 []int //
	fmt.Println("声明", sliceInt0)
	fmt.Println(len(sliceInt0))
	fmt.Println(cap(sliceInt0))

	var sliceInt02 = []int{1, 2, 3}
	fmt.Println("初始化", sliceInt02)
	fmt.Println(len(sliceInt02))
	fmt.Println(cap(sliceInt02)) //容量>=长度

}
