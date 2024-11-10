package main

import "fmt"

type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}

type Printer2 interface {
	Print()
	Print2() //必须包含结构的全部方法，才算实现了该接口
}
type Printer3 interface {
	//	空的interface  类似 Obeject 可以接收任意类型
}

func main() {
	var u user
	u.name = "Tom"
	u.age = 29
	var p Printer = u
	p.Print()

	//var p2 Printer2 = u //必须包含结构的全部方法，才算实现了该接口
	//p2.Print()

	var p3 Printer3 = u // 可以接收任意类型
	fmt.Println(p3)
}
