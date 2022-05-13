package main

import "fmt"

func main() {
	// 类似 希构函数   压栈 出栈 操作

	// return 顺序 > defer 顺序
	defer fmt.Println("End 1")
	defer fmt.Println("End 2")

	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
}
