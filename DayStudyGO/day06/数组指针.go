package main

import "fmt"

func main() {
	var array = [...]int{6, 7, 8, 9}
	fmt.Println("数组:", array)

	var arrPoint *int = &array[0] //单个值指针

	fmt.Println(arrPoint)
	fmt.Println(*arrPoint)
	//arrPoint++
	*arrPoint = *arrPoint + 4
	fmt.Println(arrPoint)
	fmt.Println(*arrPoint)
	fmt.Println("数组指针:", arrPoint)
	fmt.Println("数组:", array)

	var arrP *[4]int = &array
	fmt.Printf("数组指针 %T,%p,%d\n", arrP, arrP, arrP)
	fmt.Printf("指针取值 %T,%d\n", *arrP, *arrP)

}
