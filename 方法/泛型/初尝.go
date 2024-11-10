package main

import "fmt"

// 两种类型 泛型实例化
type Slice[T int | float32] []T

// any 作为占位符 下面支持所有类型
type AllSlice[R any] []R

// map泛型
type MyMap[KEY int | string, VALUE float64 | string] map[KEY]VALUE

//泛型结构体

// 前面泛型约束，后面是底层类型
type MyInt[T int | string] int

func main() {

	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("%T\n", a)
	var b Slice[float32] = []float32{1.1, 2.2, 3.3}
	fmt.Printf("%T\n", b)
	//类型不同，不能赋值
	//a=b
	//var c Slice[float64] =[]float64{1.0,2.0,3.0}
	var _ AllSlice[float64] = []float64{1.0, 2.0, 3.0}
	//fmt.Println(c)

	showSlice(a)
	//showSlice(b)
	//showSlice(c)

	var mp MyMap[int, string] = map[int]string{
		1: "传习录",
		2: "大问题",
	}
	fmt.Println(mp)

	//var str MyInt[string] = "12"
	var num MyInt[string] = 12
	//fmt.Println(str)
	fmt.Println(num)
}

func showSlice(s Slice[int]) {
	fmt.Println(s)
}
