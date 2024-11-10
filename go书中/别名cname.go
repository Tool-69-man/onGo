package main

import "fmt"

func main() {
	type student struct {
		name string
		id   int
	}

	s := student{id: 1, name: "土拨鼠"} // 有点像C中typeof
	fmt.Println(s)

	//	type rune = int32   可以看到rune 就是 int32 的别名
	type myStr = string
	var name myStr = "XYL"
	fmt.Println(name)
}
