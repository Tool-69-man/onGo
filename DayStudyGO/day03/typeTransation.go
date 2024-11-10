package main

import "fmt"

func main() {
	var a float32 = 12
	var b float64 = 999.123
	//a=b   //cannot use b (variable of type float64) as float32 value in assignment
	//a= (float64)b
	//a= float64(b)
	fmt.Println(a)
	fmt.Println(b)

	//var flag bool  = true
	//fmt.Printf("%T\n",flag)
	//fmt.Printf("%d\n",flag)

}
