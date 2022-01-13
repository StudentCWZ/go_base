package main

import "fmt"

// Go 语言中的运算符
func main() {
	// 1. 算术运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(5 / 2) // 2
	fmt.Println(5 % 2) // 1
	a++
	fmt.Println(a)
	// 2. 关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 <= 5)
	// 3. 逻辑运算符
	fmt.Println(10 > 5 && 1 > 0)
	fmt.Println(!(10 > 5))
	fmt.Println(1 > 5 || 3 < 4)
	// 4. 位运算符
	c := 1               // 001
	d := 5               // 101
	fmt.Println(c & d)   // 001 ==> 1
	fmt.Println(c | d)   // 101 ==> 5
	fmt.Println(a ^ b)   // 100 ==> 4
	fmt.Println(1 << 2)  // 100 ==> 4
	fmt.Println(4 >> 2)  // 100 ==> 001 ==> 1
	fmt.Println(1 << 10) // 001 ==> 10000000000 ==> 1024
	// 5. 赋值运算符
	var e int
	e = 5
	e += 5 // a = a + 5
	fmt.Println(e)

}
