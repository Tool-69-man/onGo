package main

import "fmt"

func main() {
	x := 1
	var a *int = &x
	b := &x
	//var c  *int = &x
	//var  b *float64 =&x//cannot use &x (value of type *int) as *float64 value in variable declaration

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(*a)
	//fmt.Println(b)

	fmt.Println("指针数组")
	var arr = [3]*int{a, b, &x}
	fmt.Println(arr)
	fmt.Println(*arr[1])
	//遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\t", *arr[i])
	}
	//	数组指针
	var array = [...]int{6, 7, 8, 9}
	fmt.Println("数组指针")
	//var arrPoint[len(array)]  int =array // 【】中放len(array)报错
	//arrayLen:=len(array)
	//arrayLen:=4  //这里使用也报错  GO属于静态编译

	fmt.Println("arrLen是否等于len()", 4 == len(array)) //虽然想等，但不能用

	//var arrPoint[arrayLen]  int =array//这里报错  GO属于静态编译
	//var arrPoint[arrayLen]  *int =array[0]
	var arrPoint *int = &array[0]

	fmt.Println(arrPoint)
	fmt.Println(*arrPoint)

}
