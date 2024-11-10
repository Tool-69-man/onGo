package main

import "fmt"

// (4+5)*10
func main() {
	//传入函数入参  ， 传入的函数称为回调函数，接收入参函数的函数称为高阶函数

	addAndMult(4, 5, 10, add, multFun)             //这里传入add函数变量
	r1 := multFun(add(4, 5), 10)                   //这里接收add函数的返回值
	r2 := highFun(highFun(4, 5, add), 10, multFun) //封装后，这里只调用highFun函数
	fmt.Println(r1)
	fmt.Println(r2)

}

func add(num1, num2 int) int {
	return num1 + num2
}

func multFun(x, y int) int {
	return x * y
}

func addAndMult(num1, num2, multNum int, f1, f2 func(int, int) int) {
	sum := f1(num1, num2)
	result := f2(sum, multNum)
	fmt.Println(result)
}

// 适配高阶 , 封装
func highFun(x, y int, f1 func(int, int) int) int {
	return f1(x, y)
}
