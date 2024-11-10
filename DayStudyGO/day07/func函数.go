package main

import "fmt"

// 形参 顺序，类型一一对应
func main() {

	//fmt.Println(add())
	//fmt.Println(add2(1))
	//add3(1,2,"asd")
	fmt.Println(Add4(1, 2, "asd")) //多个返回值
	fmt.Println(Sweep(1, 2))       //交换数值

}

// 重载函数：不支持
func add() int {
	return 1
}
func add2(temp int) int {
	return temp
}
func add3(temp, temp2 int, str1 string) int {
	fmt.Println("add3:", temp)
	fmt.Println(temp2)
	fmt.Println(str1)
	return temp
}
func Add4(temp, temp2 int, str1 string) (int, string) {
	return temp + temp2, str1
}
func Sweep(temp, temp2 int) (int, int) {
	return temp2, temp
}

//缺省函数不支持
//func add3(temp=1 int) int{
//	return temp
//}
