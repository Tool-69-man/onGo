package main

import "fmt"

func main() {
	//for true { //和 C 的 while true一样
	for { //for 默认为true  ，for true 一样
		var x string
		fmt.Scan(&x)
		switch x {
		case "1": //1 不做任何处理
		case "2":
			fmt.Println("2不用break")
		case "3", "4": // 3或4 起作用
			fmt.Println("逗号链接3，4")

		default:
			fmt.Println("都不是")
		}
	}

	//条件默认为true 布尔类型
	switch {
	case false:
		fmt.Println("flase")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")
	}

}
