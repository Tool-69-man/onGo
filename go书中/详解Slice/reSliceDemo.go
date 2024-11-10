package main

import "fmt"

//len = high -low
//cap = max - low

func main() {
	x := [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	showAll(x[:])     // cap 20, len 10, [0-9,外加10个0] x[0:len(slice):cap(slice)]
	showAll(x[1:2])   // cap 1,len 1, [1]  ！！cap 为19  x[1:2:20]
	showAll(x[1:2:3]) // cap 2,len 1,[1]
	showAll(x[:4])    // cap 4,len 4 ,[0-3]  ！！cap 为20  x[0:4:20]
	showAll(x[:4:7])  // cap 7,len 4,[0-3]

}

func showAll(slice []int) {
	fmt.Printf("cap: %d  len: %d  %#v\n", cap(slice), len(slice), slice)
}
