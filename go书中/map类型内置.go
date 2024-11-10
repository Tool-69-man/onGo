package main

import "fmt"

func main() {

	m := make(map[string]int) // 创建字典类型对象
	m["a"] = 1
	y, ok := m["a"]    //  ok-idiom获取值，可以知道key/value 是否存在
	fmt.Println(y, ok) // 1 true
	x, ok := m["b"]
	fmt.Println(x, ok) // 0 false

	delete(m, "a") // 删除map中a索引的值
	z, ok := m["a"]
	fmt.Println(z, ok)
}
