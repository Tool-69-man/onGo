package main

import "fmt"

func main() {
	b := 1
	fmt.Println(b, "  地址  ", &b)
	addB(&b)
	addB(&b)
	addB(&b)
	addB(&b)
	defer fmt.Println(b, "  地址  ", &b)
	defer addB(&b)
	fmt.Println(b, "  地址  ", &b)
}
func addB(b *int) {
	*b++
	fmt.Println("====", *b, &b) //多次传参，形参地址均不同
}
