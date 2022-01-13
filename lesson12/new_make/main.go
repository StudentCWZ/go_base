package main

import "fmt"

func main() {
	/*
		new 与 make 的区别
			1. 二者都是用来做内存分配的。
			2. make 只用于 slice、map 以及 channel 的初始化，返回的还是这三个引用类型本身
			3. new 用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	*/
	// a = nil
	var a *int
	a = new(int) // 初始化
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
