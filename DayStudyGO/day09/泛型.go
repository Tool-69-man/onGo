package main

import "fmt"

func main() {
	strs := []string{"爱欲与古希腊竞技", "未来日记", "图解传习录"}
	arrs := []int{1, 2, 3}
	PrintStringInterface(strs)
	fmt.Println("===")
	PrintStringArray(strs)
	fmt.Println("===")
	PrintAll(strs)
	fmt.Println("===")
	PrintAll(arrs)
}

// 法一 接口+字符数组断言
func PrintStringInterface(arr interface{}) {
	for _, v := range arr.([]string) {
		fmt.Println(v)
	}
}

// 法二 字符数组
func PrintStringArray(arr []string) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 泛型  不限定类型，让调用者限定类型   [] 定义泛型
func PrintAll[T string | int](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
