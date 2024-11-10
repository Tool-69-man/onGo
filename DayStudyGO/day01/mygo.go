package day01

import "fmt"

import _ "hello/DayStudyGO/day01/orthre" //order 依赖了 theOther
//import _ "hello/theOther"   // 这里确实不需要了
import _ "hello/DayStudyGO/day01/zzzzzzzzzz"

func main() {
	fmt.Println("HELLO MY GO")
	//var  name  = "asd"
	//fmt.Println(name)
	//
	//var  name2 string= "efg"   // 类型可写可不写 ，go自动识别
	//fmt.Println(name2)
	//
	//var  n1= 1.2
	//var  n2  float32= 1.2
	//fmt.Println(n1)
	//fmt.Println(n2)

}

func init() {
	fmt.Println("init mygo")
}
