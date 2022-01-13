package main

// 指针

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 10 // int
	b := &a // *int 类型
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, &b)
	// 变量 b 是一个 int 类型的指针(*int)，它存储的是变量 a 的内存地址
	c := *b // 根据内存地址去取值
	fmt.Println(c)
	d := 1
	modify1(d)
	fmt.Println(d)
	modify2(&d)
	fmt.Println(d)

}
