package main

import "fmt"

func main() {
	// 方式1
	var m1 map[string]string     // 声明
	m1 = make(map[string]string) // 分配空间
	m1["one"] = "java"
	fmt.Println(m1)

	// 方式2  定义
	m2 := make(map[string]string)
	m2["two"] = "C++"
	fmt.Println(m2)

	// 方式3  定义
	m3 := map[string]string{
		"one": "php",
		"two": "c++",
	}
	fmt.Println(m3)

	// 索引，通过ok来确定是否索取成功
	r, ok := m3["one"]
	fmt.Println(r, ok)

	// 使用
	m4 := map[string]string{
		"China": "Beijing",
		"USA":   "NewYork",
	}
	fmt.Println("========")
	for k, v := range m4 {
		fmt.Println(k, v)
	}
	fmt.Println("========")
	delete(m4, "China")
	for k, v := range m4 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m4 {
		fmt.Println(k)
	}
}
