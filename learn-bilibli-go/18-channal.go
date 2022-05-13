package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("发送666")
		c <- 666
	}()

	num := <-c
	fmt.Println("num = ", num)
}
