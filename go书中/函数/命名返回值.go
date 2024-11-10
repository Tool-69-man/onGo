package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(returnValue(10))
	//fmt.Println(returnValue(-111))
	fmt.Println("")
	fmt.Println(returnName(10))
	fmt.Println(returnName(-111))
}

func returnValue(a int) (int, string, error) {
	if a < 0 {
		return -1, "销售力", errors.New("自定义粗无")
	}
	return 1, "", nil
}
func returnName(a int) (num int, s string, err error) {
	if a < 0 {
		num = 2 //  num :=2 同名屏蔽 会报错  inner declaration of var num int
		s = "销售力"
		err = errors.New("故事力")
		return //隐式返回值
	}
	return
}

//cannot use -1 (untyped int constant) as string value in return statement
//显然处理return时 会跳过未命名返回值 无法准确匹配
//func returnName2(a int) (int, s string, err error) {
//	if a < 0 {
//		return -1, "销售力", errors.New("自定义粗无")
//	}
//	return 99, "", nil
//}
