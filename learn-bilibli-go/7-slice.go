package main

import "fmt"

func main() {
	// array是一个静态数组，slice是一个动态数组

	// 1. 声明一个slice
	//slice1 := []int{1,2,3}

	// 方式2,但是没有初始值
	//var slice1 []int
	//slice1 = make([]int, 3)

	// 方式3
	//var slice1 []int= make([]int, 3)

	// 方式4
	slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 2. 追加
	slice2 := make([]int, 3, 5) // 容量为5
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice2), cap(slice2), slice2)
	slice2 = append(slice2, 4)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice2), cap(slice2), slice2)

	// 3. 切片  go是指针,python是复制一份
	slice3 := []int{2, 3, 4, 5, 6}
	slice4 := slice3[1:3]
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	// 4. 拷贝
	slice5 := make([]int, 5)
	copy(slice5, slice3)
	fmt.Printf("len = %d, slice = %v\n", len(slice5), slice5)
}
