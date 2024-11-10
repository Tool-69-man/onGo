package main

import "fmt"

func main() {
	var f1 float64 = 21.21
	var f2 float32 = 21.21
	var f3 = 21.21
	f4 := 21.21
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

	fmt.Printf("%f\n", f1) //默认6位小数点
	fmt.Printf("%f\n", f2) //最好不要用float32 ,
	fmt.Printf("%f\n", f3)
	fmt.Printf("%f\n", f4)
}
