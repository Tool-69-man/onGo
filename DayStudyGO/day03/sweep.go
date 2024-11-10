package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("交换开始====")

	//a,b=b,a  //go中交换一行搞定

	//C学的方法
	sweep(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

}

func sweep(a, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}
