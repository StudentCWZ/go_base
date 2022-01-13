package main

import (
	"fmt"
	"sync"
)

// goroutine demo
var wg sync.WaitGroup // 全局

func main() { // 开启一个主 goroutine 去执行 main 函数
	wg.Add(10000) // 计数牌 + 1
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait() // 阻塞：等待所有小弟都干完活之后才收兵
}
