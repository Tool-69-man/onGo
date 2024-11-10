package main

// 定义函数类型
type FormatFunc func(s string, a ...interface{})

// 如果不用命名类型，这个参数签名长到没法看
func fotmat(f FormatFunc, s string, a ...interface{}) {
	f(s, a...)
}

func main() {

}
