package main

import "fmt"

func foo() (int, string) {
	return 10, "中国"
}

func main() {
	// 标准声明格式
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	// 批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	// 声明变量同时指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	var name2, age2 = "沙河小王子", 28
	fmt.Println(name2, age2)
	// 类型推导，编译器根据变量初始值推导出其类型
	var name3 = "小王子"
	var age3 = 18
	fmt.Println(name3, age3)
	// 短变量声明(函数内使用)
	m := 10
	fmt.Println(m)
	// 匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
}
