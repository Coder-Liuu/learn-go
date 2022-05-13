package main

import "fmt"

// 声明一种新类型
type myint int

// 定义结构体
type Book struct {
	title string
	auth  string
}

func main() {
	var b1 Book
	b1.title = "Golang"
	fmt.Println(b1)
}
