package main

import "fmt"

// range 会复制目标数据
func test(a []int) {
	for i, v := range a {
		if i == 0 {
			a[0] += 100
			a[1] += 200
			a[2] += 300
		}
		fmt.Printf("x:%d,data:%d\n", v, a[i])
	}
}
func testSlice(a []int) {
	for i, v := range a[:] {
		if i == 0 {
			a[0] += 100
			a[1] += 200
			a[2] += 300
		}
		fmt.Printf("x:%d,data:%d\n", v, a[i])
	}
}
func main() {
	data := []int{10, 20, 30}
	test(data)
	testSlice(data)
}
