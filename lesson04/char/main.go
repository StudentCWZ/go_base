package main

import "fmt"

// 字符
func main() {
	// byte	uint8 的别名 --> ASCII 码
	// rune	int32 的别名 --> 中文
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1: %T c2:%T\n", c1, c2)
	s := "hello 沙河"
	// 中文乱码
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	// 中文不乱码
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}
}
