package main

import (
	"fmt"
)

func main() {
	// Manjaro netcat(nc) 命令可以模拟tcp的链接
	fmt.Println("服务启动成功")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
