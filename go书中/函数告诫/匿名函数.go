package main

import "fmt"

func main() {

	f := func() { //
		fmt.Println("匿名函数1")
	}
	f() //如果加括号就是调用

	func() {
		fmt.Println("匿名函数2")
	}() //  加括号：调用匿名函数

	func(num, num2 int) {
		fmt.Println("可以接收入参：", num+num2)
	}(1, 2) //  匿名传入参数

	result := func(num, num2 int) int { //无返回值则无法接收
		return num + num2
	}(1, 2) //  有返回值
	fmt.Println("有返回值:", result)

}
