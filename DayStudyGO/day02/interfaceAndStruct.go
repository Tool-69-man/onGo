package main

import "fmt"

type Student struct {
	name string
}

func demo1(stu Student) {
	fmt.Println("结构体 隐式转 接口")
	demo2(stu)
}

func demo2(stu interface{}) {
	fmt.Println(stu)
}

func main() {
	//s1	:= Student{name:"熊义路"}
	//demo1(s1)
	s2 := Student{name: "熊义路"}
	demo3(s2)
}

func demo3(stu interface{}) {
	fmt.Println("接口 隐式转 结构体  编译失败")
	//demo4(stu) 失败 改用断言

	stu2 := stu.(Student)
	demo4(stu2)

	//stu3 := stu
	//demo4(stu3)
}

func demo4(stu Student) {
	fmt.Println(stu)
}
