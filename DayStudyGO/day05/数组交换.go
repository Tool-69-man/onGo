package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1
	fmt.Printf("%T,%T\n", arr1, arr2)
	fmt.Printf("%d,%d\n", arr1, arr2)

	var arr3 = [...]int{8, 7, 6}
	arr3, arr2 = arr2, arr3
	fmt.Printf("%d,%d\n", arr2, arr3)

}
