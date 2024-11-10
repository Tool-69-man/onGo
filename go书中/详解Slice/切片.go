package main

import "fmt"

// []没有声明长度，说明这是一个切片，而不是一个数组。因为数组声明是必须指定长度的。
// 切片(“动态数组”)，与数组相比切片的长度是不固定的，追加时使切片的容量增大。
func main() {
	var sliceArr = []int{1, 2, 3}
	var arr = []int{4}
	fmt.Println(sliceArr)

	fmt.Println(len(sliceArr))
	copy(sliceArr, arr) //  源码中是 []Type只能传入切面，并且类型向上转型
	//copy(sliceArr, 3)
	fmt.Println(arr)

	//sliceArr[3]=12  //错误追加
	arr2 := append(sliceArr, 12) //追加元素，但要新数组
	fmt.Println("追加后原数组", len(sliceArr))
	fmt.Println("新数组", arr2)

}
