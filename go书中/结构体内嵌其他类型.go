package main

import "fmt"

type user struct {
	name string
	age  int
}

type student struct {
	className string
	user
}

func main() {
	var s student
	s.className = "倡导动物安乐死，自己的爱犬差点离家出走"

	s.name = "戏剧与竞赛相似，都是表演"
	s.age = 22
	fmt.Println(s)
	//{倡导动物安乐死，自己的爱犬差点离家出走 {戏剧与竞赛相似，都是表演 22}}

}
