package main

import "fmt"

/**
	GO的输入
	遇到空格自动拆分
	scan  可以一行解决，也可多行
	一行解决：go自动解析拆分，根据类型，根据空格
	scanf 格式化准确无误，但编写复杂，必须完全按格式写
	scanln 规定只能一行写完，之后解析，没解析出来就是null

    注意字符串和数字好区分，但同时是字符串就解析困难
*/

func main() {
	//x:=1
	var name string
	var name2 string
	fmt.Println("请输入值")
	//fmt.Scan(&x,&name)
	//fmt.Scan(&name,&name2)
	//fmt.Println(name,"   ",name2)
	//fmt.Println("请输入值")
	////fmt.Scanf("%d %s",&x,&name) // 与c语言一样，如果逗号或空格也要写
	//
	//fmt.Scanf("%d%s",&x,&name)//这里依旧可以用空格区分,
	//// 同时也会自动区分字符串和整型
	////fmt.Scanln()
	//fmt.Println(x,"   ",name)
	//fmt.Println()

	//fmt.Println("请输入值")
	//fmt.Scanln(&x,&name)//一行读入的简写fmt.Scanf("%d%s",&x,&name)
	fmt.Scanln(&name, &name2)
	//fmt.Println(x,"   ",name)
	fmt.Println(name, "   ", name2)
}
