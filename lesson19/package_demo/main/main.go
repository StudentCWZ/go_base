package main

import (
	"fmt"
	// 当你的代码都放在 $GOPATH 目录下的话
	// 我们导入包的路径要从 $GOPATH/src 后面继续写起
	cal "lesson19/package_demo/calc" // 给导入的包起别名
)

// Go 语言中不允许导入包而不使用！！！
// Go 语言中不允许循环引用包！！！

// 多用来做一些初始化的操作
func init() {
	fmt.Println("main.init()")
}

func main() {
	fmt.Println("hello")
	sum := cal.Add(3, 5)
	fmt.Println(sum)
	fmt.Println(cal.Name)
}
