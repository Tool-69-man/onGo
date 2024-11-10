package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3}
	update(arr) //值传递 ，修改无效
	fmt.Println("实参", arr)
}

func update(arr [4]int) {
	arr[0] = 12
	fmt.Println("形参数组", arr)
}
