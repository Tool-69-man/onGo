package main

import "fmt"

func main() {
	var balance = [5]float64{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("%f", balance)
	fmt.Println(balance)

	//var arrBalancd[5] float32//这里赋值必须是5
	var arrBalancd [5]float64 //这里赋值必须是5
	arrBalancd = balance
	fmt.Println(arrBalancd)

	arr2 := [3.0]int{1, 2, 3}
	fmt.Println("长度", len(arr2))

	arr2Length := len(arr2)
	fmt.Println("数组长度", arr2Length)

	//部分初始化
	arr4 := [100]int{56: 1, 99, 1}
	fmt.Println(arr4)
	fmt.Printf("%T,%T,%b ", 4.0, 4, 4.0 == 4)

	//不定数组  也只是在定义时，下面初始化后容量是1
	//arr5 :=[...]int{1}
	//arr5=arr4

	//var arr6 =[1]int8{1}
	//	//var arr7 =[1]int16{1}
	//	//arr7=arr6
}
