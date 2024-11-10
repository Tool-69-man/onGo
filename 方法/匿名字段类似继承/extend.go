package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s student) toString() string {
	return fmt.Sprintf("stu: %+v", s)
}

type user struct {
	student //匿名嵌入其他类型
	id      int
}

func main() {
	var u user
	u.id = 1
	u.name = "窥探"
	u.age = 2
	fmt.Println(u.name) // 直接调用匿名字段成员

	fmt.Println(u.toString())
	fmt.Println(u)
	fmt.Println(fmt.Sprintf("u: %+v", u))
}
