package main

import (
	"fmt"
	out "hello/outPage"
)

//变量是一段或多段用来存储数据的内存

//我们可以通过类型转换或指针操作，
//我们可以用不同形式修改量值，
//但并不意味着改变了变量类型

func main() {
	fmt.Println(out.Myext)

	//var a , s =100,"不同类型初始化"
	var (
		x, y int
		a, s = 100, "以组的方式整理，多行变量定义"
	)

	fmt.Println(x, y, a, s)

	//简短模式
	//直接初始化
	//不提供数据类型
	//只能在函数内部
	xyl := "大问题：安德斯·尼尔森"
	fmt.Println(xyl)

	//退化赋值条件：至少有一个新变量被定义，且必须是同一作用域
	xyl, y2 := "退化赋值", 666
	fmt.Println("同时需要赋值和初始化：", y2, xyl)

	//退化赋值 允许我们重复使用 同一个变量名
	// 比如 err , ok

	{
		xyl := "不同作用域，全新变量定义"
		fmt.Println(xyl)
	}

}
