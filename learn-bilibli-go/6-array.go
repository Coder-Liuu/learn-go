package main

import "fmt"

func main() {
	// 定义数组
	fmt.Println("a1")
	var a1 [3]int
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}

	// 定义方法2 + 初始化
	fmt.Println("\na2")
	a2 := [3]int{1, 2, 3}
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}

	// 定义方法2 + 遍历
	fmt.Println("\na3")
	a3 := [3]int{99, 19, 9}
	for index, value := range a3 {
		fmt.Println("index = ", index, "value = ", value)
	}
}
