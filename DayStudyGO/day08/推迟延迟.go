package main

import "fmt"

// 一个方法延迟 推迟
// 栈的方式执行 后进先出  资源的释放，比如关闭io操作
//
//	指针不起作用，推迟的a 一直是1
//
// 延迟函数的值都打印 a=1,
func main() {
	a := 1

	defer fmt.Println("第", 4, "延迟", a, &a) //结果不变 还是1
	fmt.Println("最先显示", a, &a)
	defer addA(&a, 3)                      //   推迟传入的a为啥地址不同
	defer fmt.Println("第", 3, "延迟", a, &a) //
	defer addA(&a, 2)                      //
	defer fmt.Println("第", 2, "延迟", a, &a) //
	defer addA(&a, 1)                      //结果为2
	defer fmt.Println("第", 1, "延迟", a, &a) //
	fmt.Println("最后一行,值不变", a, &a)
}

func addA(a *int, num int) {
	*a += 1
	fmt.Println("后进先出延迟:", num, " 先一步显示", *a, &a)
}
