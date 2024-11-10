package main

import "fmt"

func main() {
	for i, c := 0, count(); i < c; i++ { //count 只执行一次
		fmt.Println(i)
	}

	for i := 0; i < count(); i++ {
		fmt.Println("重复count", i)
	}

	for i := 0; i < count(); i++ {

	}

	//	冒泡排序
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//fmt.Println(arr)

			}
		}
	}

	fmt.Println()

}
func count() int {
	fmt.Println("count======")
	return 3
}
