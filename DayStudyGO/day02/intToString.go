package main

import "fmt"

func main() {

	num := 2.1
	fmt.Printf("%T \n", num)
	var i int = int(num)
	fmt.Printf("%T \n", i)

	num2 := 97
	var i2 string = string(num2)
	fmt.Printf("%T \n", i2)
	fmt.Printf("%s \n", i2)

	num3 := (string)(65) //这里前面多个括号
	fmt.Printf("%T \n", num3)
	fmt.Printf(num3)

}
