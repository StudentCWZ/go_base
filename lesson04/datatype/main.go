package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%d\n", n) // 十进制
	fmt.Printf("%b\n", n) // 二进制
	// 八进制
	m := 075
	fmt.Printf("%d\n", m) // 十进制
	fmt.Printf("%o\n", m) // 八进制
	// 十六进制
	f := 0xff
	fmt.Println(f)        // 十进制
	fmt.Printf("%x\n", f) // 十六进制
	// unit8
	var age uint8
	// age = 256	// 数值溢出
	fmt.Println(age)
	// 浮点数(最大值)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	// 布尔值
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	// 字符串
	s1 := "hello beijing"
	s2 := "你好 北京"
	fmt.Println(s1)
	fmt.Println(s2)
	// 打印 windows 平台下的一个路径 c:\code\go.exe
	fmt.Println("c:\\code\\go.exe")
	fmt.Printf("\t制表符\n换行符\n")
	s3 := `
	多行字符串
	两个反引号之间的内容
	回原样输出
	\t
	\n
	`
	fmt.Println(s3)
}
