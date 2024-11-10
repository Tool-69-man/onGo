package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	b := "abcdefg"
	b1 := b[2:5]
	fmt.Println(b, b1)
	fmt.Println(unsafe.Pointer(&b))
	fmt.Printf("%#v\n", unsafe.Pointer(&b))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&b)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&b1)))
	//	 %#  能显示结构
	//   unsafe.Pointer 用于指针类型转换
}
