package main

import "fmt"

type sliceDemo[T int | float32] []T

func main() {
	var a sliceDemo[int] = []int{1, 2, 3}
	fmt.Printf("%T\n", a)
	var b sliceDemo[float32] = []float32{1.1, 2.2, 3.3}
	fmt.Printf("%T\n", b)
	// 方法
	a.showSlice()
	b.showSlice()

}

func (s sliceDemo[T]) showSlice() {
	fmt.Println(s)
}
