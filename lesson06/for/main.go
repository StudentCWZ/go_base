package main

import "fmt"

// for 循环
func main() {
	// 1. for 循环的基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 2. 省略初始语句，但是必须保留初始语句后面的分号
	var i = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	// 3.省略初始语句和结束语句
	var u = 10
	for u > 0 {
		fmt.Println(u)
		u--
	}
	// 4. 无限循环
	// for {
	// 	fmt.Println("Hello 沙河！")
	// }
	// 5. break 跳出 for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
	// 6. continue 继续下一次循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // 跳过本次 for 循环，继续下一次循环
		}
		fmt.Println(i)
	}
	/*
		7. for ... range(后面再详细介绍): Go 语言中可以使用 for range 遍历数组、切片、字符串、map 及通道(channel)，
			通过 for range 遍历的返回值有以下规律：
				(1) 数组、切片、字符串返回索引和值
				(2) map 返回键和值
				(3) 通道(channel) 只返回通道内的值
	*/
}
