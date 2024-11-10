package main

import "fmt"

func main() {
	v := 1
	v++ //无 前缀 ++v
	//fmt.Printf("%d", v++) 不能同步写
	// syntax error: unexpected ++ in argument list; possibly missing comma or )
	fmt.Printf("%d\n", v)

	if v == 2 {
		v--
		fmt.Println(v)
	} else if v > 2 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}
