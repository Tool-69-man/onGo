package main

import "fmt"

func main() {
	sum(1, 2, 3)
}

func sum(arg ...int) {
	var result int
	for e := range arg {
		result += arg[e]
	}
	fmt.Println(result)
}
func sum2(arg ...int) {
	var result int
	for e := range arg {
		result += arg[e]
	}
	fmt.Println(result)
}
