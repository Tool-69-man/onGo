package main

import (
	"fmt"
)

func main() {
	//i := 1
	for i := 1; i < 10; i++ { //不好用 var 为啥？
		fmt.Println(i)
	}

	var a [10]int //数组声明

	for i := 0; i < 10; i++ { //不好用 var 为啥？
		a[i] = i + 1
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", a[i])
	}

}
