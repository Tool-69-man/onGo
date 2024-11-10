package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	update(arr)
	fmt.Println("实参", arr)
}

//去掉数组范围，是go中切片slice
func update(arr []int) { //切片说引用类型，入参直接是地址
	arr[0] = 12
	arr[2] = 12
	fmt.Println("形参数组", arr)
}
