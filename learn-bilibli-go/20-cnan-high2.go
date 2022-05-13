package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		// 如果不关闭会 死锁
	}()

	//for {
	//data, ok := <-c
	//if ok {
	//fmt.Println("data = ",data)
	//} else {
	//break
	//}
	//}

	for data := range c {
		fmt.Println("data = ", data)
	}

	fmt.Println("Main End")
}
