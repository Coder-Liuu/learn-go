package main

import "fmt"

// 基础知识：for循环
func main() {
	for {
		fmt.Println("Loop")
		break
	}

	for i := 0; i < 3; i++ {
		fmt.Println("i = ", i)
	}

	for val, idx := range []int{1, 2, 3, 4} {
		fmt.Println("val=", val, "idx=", idx)
	}
}
