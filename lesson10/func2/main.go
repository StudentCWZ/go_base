package main

import "fmt"

// 函数进阶之变量作用域
// 定义全局变量
var num = 10

// 定义一个函数
func testGlobal() {
	num := 100
	name := "沙河小王子"
	// 函数中使用变量
	// 1. 先在自己函数中查找，找到了就用自己函数中的
	// 2. 函数中找不到变量就往外层找全局变量
	fmt.Println("变量 num:", num)
	fmt.Println(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	fmt.Println("变量 num:", num)
	testGlobal()
	// 外层不能访问到函数的内部变量(局部变量)
	// fmt.Println(name)
	// 变量 i 此时只在 for 循环的语句块中生效
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// 外层访问不到内部 for 语句块中的变量
	// fmt.Println(i)
	// 函数是可以作为变量
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()
	r := cal(100, 200, add)
	fmt.Println(r)
	r1 := cal(100, 200, sub)
	fmt.Println(r1)
}
