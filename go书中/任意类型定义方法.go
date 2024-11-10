package main

import "fmt"

type X int

func (x *X) inc() { //名称前参数称作 receiver 作用类似 python self
	*x++
}

func main() {

	var x X
	x.inc()
	fmt.Println(x)

}
