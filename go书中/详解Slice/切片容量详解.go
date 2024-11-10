package main

import "fmt"

func main() {
	//	切片相当于 写字楼
	//  长度代表正在用的楼层
	//  容量 最高是栋楼高，容量必须大于长度，并且地址连续
	//  每层能容100人，你公司200人，但公司未来会发展到700人。
	//  你去和租写字楼的老板说：我先买两层，未来可以还会再买4，5层。能先给我留着吗？

	company := make([]int, 2, 7)
	fmt.Printf("买来的公司:%d,位置%p\n", company, company)
	fmt.Println(len(company))
	fmt.Println(cap(company))

	//company.rang
	for e := range company {
		company[e] += 1
	}
	fmt.Printf("公司迅速发展:%d,位置%p\n", company, company)
	fmt.Println(len(company))
	fmt.Println(cap(company))

	company = append(company, 1) // 公司人数增加，在原来基础再买一层
	fmt.Printf("公司扩展一层：%d，地址无变化：%p\n", company, company)
	fmt.Println(len(company))
	fmt.Println("长度无需扩容：", cap(company))

	for i := 2; i < 10; i++ {
		company = append(company, i)
	}
	fmt.Printf("新的公司，更多空间扩展：%d，新位置%p\n", company, company)
	fmt.Println("公司1100人：", len(company))
	fmt.Println("容量翻倍增加", cap(company))

	//容量的扩展规律是按容量的 2 倍数进行扩充

	company = append(company, 9, 8, 7, 6, 5, 4) //多元素添加
	fmt.Printf("进一步扩容：%d，位置%p\n", company, company)
	fmt.Println(len(company))
	fmt.Println(cap(company))

	company2 := []int{1, 2, 3}
	company = append(company, company2...) //追加一个切片, 切片需要解包 ，slice 后加 '...'
	fmt.Printf("吞并一个小公司：%d，位置%p\n", company, company)
	fmt.Println(len(company))
	fmt.Println(cap(company))

	//第一个入参必须为切片
	company = append([]int{999}, company...) //切片前添加一个元素（将切面元素，添加到别的切面后面），
	fmt.Printf("公司被政府收购：%d，位置%p\n", company, company)
	fmt.Println(len(company))
	fmt.Println("成了新的切片，容量自动缩小", cap(company))

}
