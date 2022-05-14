package main

import (
	"fmt"
	"time"
)

// 基础知识：switch
func main() {
	var a = 2
	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	default:
		fmt.Println("a = unknow")
	}

	var t = time.Now()
	// 这里不需要写t，可以替代if语句
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	case t.Hour() >= 12:
		fmt.Println("lt's after noon")
	}
}
