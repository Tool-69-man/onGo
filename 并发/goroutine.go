package main

import (
	"fmt"
	"time"
)

// go
func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, "数值：", i)
		time.Sleep(time.Second)
	}
}

func main() {
	//在task前 加go  则新建一个线程
	//注意 要让main 线程延长时间结束，否则没等task，开始main就结束了
	go task("任务1")
	task("任务2")

	time.Sleep(time.Second)
}
