package main

import "fmt"

// 结构体的初始化
type person struct {
	name, city string
	age        int8
}

// 结构体占用一块连续的内存
type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	// 1. 键值对初始化方式
	p := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p: %#v\n", p)
	p1 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p1: %#v\n", p1)
	// 2. 值的列表进行初始化
	/*
		使用值的列表初始化时，需要注意：
			1. 必须初始化结构体的所有字段
			2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致
			3. 该方式不能和键值初始化方式混用
	*/
	p2 := person{
		"小王子",
		"北京",
		18,
	}
	fmt.Printf("p2: %#v\n", p2)
	// 结构体占用一块连续的内存
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}
