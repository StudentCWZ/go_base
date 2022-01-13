package main

import "fmt"

// 定义结构体
// type person struct {
// 	name string
// 	age int
// 	city string
// }

// 简写
type person struct {
	name, city string
	age        int8
}

func main() {
	var p person
	p.name = "沙河娜扎"
	p.city = "北京"
	p.age = 18
	fmt.Printf("p = %#v\n", p)

	// 匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "小王子"
	user.married = false
	fmt.Printf("user = %#v\n", user)
}
