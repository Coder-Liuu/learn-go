package main

import (
	"fmt"
	"time"
)

func main() {
	// 有缓冲的channal
	// 当channal满或者空，阻塞
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("goroutine 结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go正在运行,发送元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("End")

	time.Sleep(2 * time.Second)
}
