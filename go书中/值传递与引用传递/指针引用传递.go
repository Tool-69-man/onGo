package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	update2(&arr)
	fmt.Println("实参", arr)
}

func update2(arr *[3]int) {
	(*arr)[1] = 12
	arr[0] = 12
	arr[2] = 12
	fmt.Println("形参数组", *arr)
}
