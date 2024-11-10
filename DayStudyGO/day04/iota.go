package main

import "fmt"

const (
	a1 = 2
	a2 = 2
	a3 = iota
)

const (
	b1 = iota
	b2 = iota
	b3
)

const c1, c2, c3 = iota, iota, iota

//const name1 =iota
//const name2
//const name3 =iota
//const name4

func main() {
	fmt.Println(a1 == a2, a1 == a3, a2 == a3)
	fmt.Println(a1, " ", a2, " ", a3)
	fmt.Println(b1, " ", b2, " ", b3)
	fmt.Println(c1, " ", c2, " ", c3) //都是0，全新iota
	//fmt.Println(name1)
	//fmt.Println(name2)
	//fmt.Println(name3)
	//fmt.Println(name4)
}
