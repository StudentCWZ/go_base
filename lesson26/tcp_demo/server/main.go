package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 处理完要关闭这个连接
	// 针对当前连接做数据的发送和接收操作
	for {
		read := bufio.NewReader(conn)
		var buf [128]byte
		n, err := read.Read(buf[:])
		if err != nil {
			fmt.Printf("Read from conn failed, err: %v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据:", recv)
		conn.Write([]byte("ok")) // 把数据返回客户端
	}

}

// tcp server demo

func main() {
	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Listen failed, err: %v\n", err)
		return
	}
	for {
		// 2. 等待客户端来建立连接
		conn, err := listen.Accept() // 阻塞
		if err != nil {
			fmt.Printf("Accept failed, err: %v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
}
