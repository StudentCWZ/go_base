package main

import (
	"fmt"
	"sync"
)

// sync.map 并发安全的 map
var (
	// m  = make(map[int]int)
	wg sync.WaitGroup
	m2 = sync.Map{}
)

// func get(key int) int {
// 	return m[key]
// }

// func set(key int, value int) {
// 	m[key] = value
// }

func main() {
	// m 版本
	// for i := 0; i < 20; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		set(i, i+100)                             // map 中设置键值对
	// 		fmt.Printf("key: %v value:%v", i, get(i)) // 打印键值对
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// m2 版本
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key: %v value:%v\n", i, value) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
