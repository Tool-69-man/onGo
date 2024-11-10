package main

//库源码文件： 不含启动类 ， 不能在main包下
//源码文件： 程序入口，必须在main包下

import (
	"fmt"
	util "hello/go书中/工具包" // 可以给库源添加前缀 ,库资源包
	//. "hello/工具包" // 可以给库源添加前缀 ,库资源包
	. "hello/go书中/库源码文件包" // 可以用点代替前缀,简易导包不允许有函数名重复
)

func main() {
	fmt.Println(util.Add4(1, 2, "asd")) // 资源包 导出函数 首字母大写
	//fmt.Println(Add4(1, 2, "asd")) // 简易导包不允许有函数名重复
	fmt.Println(Sweep(1, 2)) // 这里属于day07导出的函数
}
