package main

import "fmt"

func main() {
	v1 := 'A' //单引号 类型直接是整型 int32,能用%c接收
	v2 := "A"

	fmt.Printf("%s\n", v1) //%!s(int32=65)

	fmt.Printf("%s\n", v2)

	fmt.Printf("%T,%c\n", v1, v1) //int32  ,  A
	fmt.Printf("%T,%d\n", v1, v1) //int32 , 65
}
