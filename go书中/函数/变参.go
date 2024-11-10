package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	x2 := [4]int{1, 2, 3, 4}

	test(12, 23, 4)
	test(x...)     //slice 传入
	test(x2[:]...) //数组reSlice后展开
}

// 变参就是 切片，放在列表尾部
func test(a ...int) {
	fmt.Println(a)
}
