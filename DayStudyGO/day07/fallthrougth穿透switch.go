package main

import (
	"fmt"
)

// switch  不用再写break 却增加了fallthrough
// 无条件进入下一个case
func main() {
	var exp2 int = -11
	x := 1
	y := 2
	z := 3
	switch exp2 {
	default:
		fmt.Println("放在第一位也没事")
		fallthrough
	case x:
		fmt.Println("x中显示", x)
		fallthrough // 进入下一个case
		//fmt.Println("x中显示",x)   //fallthrough 下面不能写语句
	case y:
		fmt.Println("y中显示", y) //
		fallthrough
	case z:
		// fallthrougth 条件不符合我的要求，就应该推出
		if exp2 < (-1) {
			fmt.Println("条件不符合：", exp2)
			break
		}
		fmt.Println("z")
	}

}
