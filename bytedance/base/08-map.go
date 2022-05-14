package main

import "fmt"

// 基础知识：字典的使用
func main() {
	var m = make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["two"])    // 2
	fmt.Println(m["unknow"]) // 0

	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")

	var m2 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2)
}
