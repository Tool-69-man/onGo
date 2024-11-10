package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	x := 0
	//if x,err:=10,check(x);err == nil {  x与check(x)不能同时初始化,必须使用已初始化的值
	if err := check(x); err == nil { //
		log.Fatalln("数字正确")
	} else {
		fmt.Println(err)
	}
	fmt.Println("END")
}
func check(x int) error {
	if x <= 0 {
		return errors.New("自定义错误：x<=0")
	}
	return nil
}
