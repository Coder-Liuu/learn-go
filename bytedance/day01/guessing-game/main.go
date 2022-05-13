package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用时间戳作为随机数种子
	rand.Seed(time.Now().UnixNano())

	maxNum := 100
	var myNum int

	secretNumber := rand.Intn(maxNum)
	fmt.Println("secretNumber is", secretNumber)
	for {
		fmt.Scanln(&myNum)
		if myNum > secretNumber {
			fmt.Println("太大了")
		} else if myNum < secretNumber {
			fmt.Println("太小了")
		} else {
			fmt.Println("恭喜你猜对了,exit(0)")
			break
		}
	}
}
