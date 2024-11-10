package main

import "fmt"

func main() {
	//能直接以数组的索引取值
	var str string = "xiong yi lu" // 每个字符 是uint8
	fmt.Println(str, " 长度：", len(str))
	fmt.Println(str[0], str[1], str[2])
	fmt.Printf("%s%s%s\n", str[0], str[6], str[9])          //%!s(uint8=120)%!s(uint8=121)%!s(uint8=108)
	fmt.Printf("%c%c%c\n", str[0], str[6], str[9])          //结果 xyl
	fmt.Printf("%c%c%c\n", str[0]-32, str[6]-32, str[9]-32) //结果 xyl
	//println(str)//标红输出

	//升级版 for循环  range
	for e := range str {
		fmt.Printf("索引：%3d  ASCII:%3d  字符：%c\n", e, str[e], str[e])
	}
	fmt.Println()
	for k, v := range str {
		fmt.Printf("索引：%3d  ASCII:%3d  字符：%c\n", k, v, v)
	}
	fmt.Println()
	p := &str
	fmt.Printf("%d %p %s\n", p, p, *p)
	fmt.Printf("%c %d\n", (*p)[0], (*p)[0]+4)

	//str[0]='x'  //str字符不能再修改
	fmt.Println(str[0])
}
