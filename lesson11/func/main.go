package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
// 定义一个函数它的返回值是一个函数
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 赋给变量，然后执行
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()
	// 直接执行
	func() {
		fmt.Println("匿名函数")
	}()
	name := "沙河小王子"
	// 闭包 = 函数 + 外层变量的引用
	r := a(name) // r 此时就是一个闭包
	r()          // 相当于执行了 a 函数内部的匿名函数
	// 闭包
	r1 := makeSuffixFunc(".txt")
	ret := r1("沙河小王子")
	fmt.Println(ret)
	// 闭包
	x, y := cal(100)
	add := x(200)
	sub := y(200)
	fmt.Println(add, sub)
}
