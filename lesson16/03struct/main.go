package main

import "fmt"

// 嵌套结构体的字段冲突
type Address struct {
	Province   string
	City       string
	UpdateTime string
}

type Email struct {
	Addr       string
	UpdateTime string
}

type Person struct {
	Name    string
	Gender  string
	Age     int
	Address // 嵌套另外一个结构体
	Email   // 嵌套另一个结构体
}

func main() {
	p1 := Person{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province:   "山东",
			City:       "威海",
			UpdateTime: "2019-07-11",
		},
		Email: Email{
			Addr:       "小王子@163.com",
			UpdateTime: "2018-07-01",
		},
	}
	fmt.Printf("p1: %#v\n", p1)
	// fmt.Println(p1.UpdateTime) 报错：嵌套结构体中包含同名的 UpdateTime
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}
