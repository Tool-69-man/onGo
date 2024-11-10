package main

import "fmt"

func main() {
	temp := 1
	fmt.Println(temp)
	if true {
		temp = 3
		temp := 2
		fmt.Println(temp)
		temp = 4
	}
	fmt.Println(temp) // 全局变量
}
