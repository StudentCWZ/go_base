package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与 json 序列化
// 如果一个 Go 语言包中定义的标识符是首字母大写的，那么就是对外可见的
// 如果一个结构体中的字段名首字母是大写的，那么该字段就是对外可见的
type student struct {
	ID   int
	Name string
}

// student 的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title    string    `json:"title"`
	Students []student `json:"students"`
}

func main() {
	// 创建一个班级变量 c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	// 往班级 c1 中添加学生
	for i := 0; i < 10; i++ {
		// 十个学生
		temStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, temStu)
	}
	fmt.Printf("c1: %#v\n", c1)
	// JSON 序列化：Go 语言中的数据 -> JSON 格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Printf("json marshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	// JSON 反序列化：JSON 格式的字符串 -> Go 语言的数据
	jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Printf("json unmarshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("%T\n", c2)
	fmt.Printf("%#v\n", c2)
}
