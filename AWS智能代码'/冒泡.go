package main

import "fmt"

func main() {
	//	冒泡排序
	arr := []int{1, 5, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(arr)
			}
		}
	}
}
