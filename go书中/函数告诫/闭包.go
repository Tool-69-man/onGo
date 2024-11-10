package main

import "fmt"

// 内层函数中操作外层函数变量，内层函数和外层函数统称为闭包结构

func main() {
	f := increment2()
	result := f()
	fmt.Println(result, &result)

	result = f()
	fmt.Println(result, &result)
	result = f()
	fmt.Println(result, &result)
	result = f()
	fmt.Println(result, &result)

	f2 := increment2()
	result2 := f2()
	fmt.Println(result2, &result2)

	result2 = f2()
	fmt.Println(result2, &result2)
	result2 = f2()
	fmt.Println(result2, &result2)
	result2 = f2()
	fmt.Println(result2, &result2)

}

// 无参版本
func increment() func() {
	i := 0
	f := func() {
		i++
	}
	return f
}
func increment2() func() int {
	i := 0
	//局部变量
	f := func() int {
		i++
		return i
	}
	return f
}
