package main

import (
	"fmt"
	"sync"
)

// goroutine demo
var wg sync.WaitGroup // 全局

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done() // 通知 wg 把计算器 - 1
}

func main() { // 开启一个主 goroutine 去执行 main 函数
	wg.Add(10000) // 计数牌 + 1
	for i := 0; i < 10000; i++ {
		// wg.Add(1)
		go hello(i) // 开启了一个 goroutine 去执行 hello 这个函数
	}
	fmt.Println("hello main")
	// time.Sleep(time.Second)
	wg.Wait() // 阻塞：等待所有小弟都干完活之后才收兵
}
