package main

import "fmt"

// 定义常量
// iota默认是1  并且只能配合const使用
const (
	A = iota + 1
	B
	C
)

func main() {
	fmt.Println("A = ", A)
	fmt.Println("B = ", B)
	fmt.Println("C = ", C)
}
