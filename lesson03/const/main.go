package main

import "fmt"

// 常量
// const pi = 3.1415
// const e = 2.7

// 批量声明常量
// const (
// 	pi = 3.1415
// 	e = 2.7
// )

// 批量声明常量
// const (
// 	n1 = 10
// 	n2
// 	n3
// )

// const (
// 	n1 = iota // 0
// 	n2        // 1
// 	n3        // 2
// 	n4        // 3
// )

// 单下划线跳过值
// const (
// 	n1 = iota // 0
// 	n2        // 1
// 	_         // 2
// 	n4        // 3
// )

// iota 中间插队
const (
	n1 = iota // 0
	n2 = iota // 1
	n3 = 100
	n4 = iota // 3
)
const n5 = 100

// 定义数量级
// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB = 1 << (10 * iota)
// 	GB = 1 << (10 * iota)
// 	TB = 1 << (10 * iota)
// 	PB = 1 << (10 * iota)
// )

const (
	a, b = iota + 1, iota + 2 // 1, 2
	c, d                      // 2, 3
	e, f                      // 3, 4
)

func main() {
	// fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(n1, n2, n3, n4, n5)
	fmt.Println(a, b, c, d, e, f)
}
