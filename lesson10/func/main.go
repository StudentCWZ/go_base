package main

import "fmt"

// 定义一个不需要参数也没有返回值的函数
// func sayHello() {
// 	fmt.Println("Hello 沙河小王子！")
// }

// 定义一个接收 string 类型的 name 参数
func sayHello(name string) {
	fmt.Println("Hello", name)
}

// 定义接收多个参数的函数
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSumTwo(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数
func intSumThree(a ...int) int {
	ret := 0
	for _, args := range a {
		ret = ret + args
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSumFour(a int, b ...int) int {
	ret := a
	for _, args := range b {
		ret = ret + args
	}
	return ret
}

// Go 语言中函数参数类型的间歇
func intSumFive(a, b int) (ret int) {
	ret = a + b
	return
}

// 定义具有多个返回值的函数
func sumAndSub(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	name := "沙河小王子"
	// 函数调用
	sayHello(name)
	r := intSum(5, 10)
	fmt.Println(r)
	r2 := intSumTwo(10, 20)
	fmt.Println(r2)
	r3 := intSumThree(10, 20)
	fmt.Println(r3)
	r4 := intSumFour(10, 20, 10) // a = 10, b = [20, 10]
	fmt.Println(r4)
	r5 := intSumFive(10, 20)
	fmt.Println(r5)
	x, y := sumAndSub(100, 200)
	fmt.Println(x, y)
}
