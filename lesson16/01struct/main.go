package main

import "fmt"

type Person struct {
	string
	int
}

// 结构体的匿名字段
func main() {
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}
