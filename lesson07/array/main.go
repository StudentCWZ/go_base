package main

import (
	"fmt"
	"sort"
)

func f1(a [3][2]int) {
	a[0][0] = 100
}

// 数组相关内容
func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[0])
	fmt.Println(cityArray[3])
	// 2. 编译器推导数组的长度
	var boolArray = [...]bool{true, false, false, true, false}
	fmt.Println(boolArray)
	// 3. 使用索引值方式初始化
	var langArray = [...]string{1: "Golang", 3: "Python", 7: "java"}
	fmt.Println(langArray)
	fmt.Printf("%T\n", langArray)
	// 数组的遍历
	// 1. 使用 for 循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	// 2. 使用 for range 方式遍历
	for index, value := range cityArray {
		fmt.Println(index, value)
	}
	// 二维数组
	cityArrayTwo := [4][2]string{
		{"北京", "西安"},
		{"上海", "重庆"},
		{"杭州", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(cityArrayTwo)
	// 重庆
	fmt.Println(cityArrayTwo[1][1])
	// 二维数组的遍历
	for _, v1 := range cityArrayTwo {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组是值类型
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	// 值拷贝
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	// 值拷贝
	y := x
	y[0][0] = 100
	fmt.Println(x)
	// 练习题
	// 请使用内置的 sort 包对数组 var w = [...]int{3, 7, 8, 9, 1} 进行排序
	var w = [...]int{3, 7, 8, 9, 1}
	// w[:] 得到的是一个切片，指向了底层的数组 w
	sort.Ints(w[:])
	fmt.Println(w)
}
