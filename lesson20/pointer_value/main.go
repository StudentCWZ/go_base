package main

import "fmt"

// 使用值接收者实现接口和使用指针接收者实现接口的区别

type mover interface {
	move()
}

type sayer interface {
	say()
}

// 接口嵌套
type animal interface {
	sayer
	mover
}

type person struct {
	name string
	age  int8
}

// 使用值接收者实现接口：类型的值和类型的指针都能保存到接口变量中
// func (p person) move() {
// 	fmt.Printf("%s在跑...\n", p.name)
// }

// 使用指针接收者实现接口：只有类型指针能保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s谁在叫～\n", p.name)
}

func main() {
	var m mover // 定义一个 mover 类型的变量
	// p1 := person{ // p1 是 person 类型的值
	// 	name: "小王子",
	// 	age:  18,
	// }
	p2 := &person{ // p2 是 person 类型的指针
		name: "娜扎",
		age:  18,
	}
	// m = p1 // 报错：无法复制，因为 p1 是 person 类型的值没有实现 mover 接口
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer // 定义一个 sayer 类型的变量
	s = p2
	s.say()
	fmt.Println(s)

	var a animal // 定义一个 animal 类型的变量
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
