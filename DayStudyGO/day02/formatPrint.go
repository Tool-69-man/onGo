package main

import "fmt"

func main() {
	var url string = "age=%d And show:%s"
	//fmt.Println(url,"22","如此反复怎成大气")
	var result = fmt.Sprintf(url, 22, "如此反复怎成大气")
	fmt.Println(result)
}
