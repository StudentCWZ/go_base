package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// map(映射)
func main() {
	// 光声明 map 类型，但是没有初始化，a 就是初始值 nil
	var a map[string]int
	fmt.Println(a == nil)
	// map 的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	// map 如何添加键值对
	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200
	fmt.Println(a)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("%T\n", a)
	// 声明 map 的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b: %#v\n", b)
	fmt.Printf("%T\n", b)
	var c map[int]int
	// c[100] = 200 // 报错：这个 map 没有初始化
	fmt.Println(c)
	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["沙河娜扎"] = 100
	scoreMap["沙河小王子"] = 200
	// 判断张二狗子在不在 scoreMap 中
	v, ok := scoreMap["张二狗子"]
	fmt.Println(v, ok)
	if !ok {
		fmt.Println("查无此人!")
	} else {
		fmt.Println("张二狗子在 scoreMap 中", v)
	}
	// 遍历
	// map 是无序的，键值对和添加的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 删除键值对
	delete(scoreMap, "沙河小王子")
	fmt.Println(scoreMap)
	// 按照某个固定顺序遍历 map
	var intMap = make(map[string]int, 100)
	// 添加 50 个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0~99 的随机整数
		intMap[key] = value
	}
	for k, v := range intMap {
		fmt.Println(k, v)
	}
	// 按照 key 从小到大的顺序去遍历 intMap
	// 1. 先取出所有的 key 存放到切片中
	keys := make([]string, 0, 100)
	for k := range intMap {
		keys = append(keys, k)
	}
	// 2. 对 key 排序
	sort.Strings(keys)
	// 3. 按照排序后的 key 对 intMap 排序
	for _, key := range keys {
		fmt.Println(key, intMap[key])
	}
	// 元素类型为 map 的切片
	var mapSlice = make([]map[string]int, 8) // 只是完成了切片的初始化
	// [nil nil nil nil nil nil nil nil]
	fmt.Println(mapSlice[0] == nil) // true
	// mapSlice[0]["沙河小王子"] = 100               // 报错: assignment to entry in nil map
	// 还需要完成内部 map 元素的初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["沙河小王子"] = 100
	fmt.Println(mapSlice)
	// 值为切片的 map
	var sliceMap = make(map[string][]int, 8) // 只完成 map 初始化
	value, ok := sliceMap["中国"]
	if ok {
		fmt.Println(value)
	} else {
		// slicemap 中没有中国这个键
		sliceMap["中国"] = make([]int, 8) // 完成对切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][3] = 300
	}
	// 遍历 sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
	// 练习题
	// 统计一个字符串中每个单词出现的次数
	// "how do you do" 中每个单词出现的次数
	// 1. 定义一个 map[string][int]
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	// 2. 字符串中都有哪些单词
	words := strings.Split(s, " ")
	// 3. 遍历单词做统计
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			// map 有这个单词的统计记录
			wordCount[word] = v + 1
		} else {
			// map 没有这个单词的统计记录
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
