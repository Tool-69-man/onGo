package main

import "fmt"

func main() {
	fmt.Printf("%T\n", f1) // 类型 为func() , 函数也成了类型 ，可以定义函数变量
	fmt.Printf("%d\n", f1) // f1不带括号就是变量
	fmt.Printf("%p\n", f1) //  带了括号就成了函数调用
	fmt.Println(f1)

	fmt.Println("\n函数变量\n")

	var f2 func(string, string)
	f2 = f1

	f3 := f1
	f1("行而突破表现", "而享其道")
	f2("图书馆志愿者", "")
	f3("内心强大", "不如实现意义的人")

	fmt.Println("3个函数内存地址相同：", f1, "  ", f2, "  ", f3) //本质上还是调用f1

}
func f1(str, str2 string) {
	fmt.Println(str)
	fmt.Println(str2)
}
