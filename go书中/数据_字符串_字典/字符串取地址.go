package main

import "fmt"

func main() {
	a := "abcd"

	fmt.Println(a[0])
	//fmt.Println(&a[1])//不能获取地址‘

	//	 切片取字符串

	fmt.Println(a[:])
	fmt.Println(a[1:1]) //空
	fmt.Println(a[1:2]) //b
	fmt.Println(a[:3])  //abc
	fmt.Println(a[3:])  //d

}
