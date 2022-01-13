package main

import "fmt"

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// 构造函数
// NewPerson 是一个 Person 类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 是为 Person 类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好 Go 语言!\n", p.name)
}

// setAge 是一个修改年龄的方法
// 指针的接收者指的就是接收者的类型是指针
func (p *Person) setAge(age int8) {
	p.age = age
}

// setAge2 是一个使用值接收者来修改年龄的方法
func (p Person) setAge2(age int8) {
	p.age = age
}

// 方法的定义实例
func main() {
	p := NewPerson("沙河娜扎", int8(18))
	(*p).Dream()
	p.Dream()
	// 调用修改年龄的方法
	fmt.Println(p.age) // 18
	p.setAge(int8(20))
	fmt.Println(p.age) // 20
	p.setAge2(int8(30))
	// 值拷贝
	fmt.Println(p.age) // 20
	/*
		什么时候应该使用指针类型？
			1. 需要修改接收者中的值
			2. 接收者是拷贝代价比较大的对象
			3. 保持一致性，如果某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	*/

}
