package main

import "fmt"

// 基础知识：切片
func main() {
	var s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get 0", s[0])
	fmt.Println("len", len(s))

	// 注意append的用法的话，必须把 append的结果赋值为原数组。
	s = append(s, "e", "f")
	fmt.Println("s", s)
	fmt.Println("len", len(s))

	// 修改原数组，不会影响新数组
	s2 := make([]string, len(s))
	copy(s2, s)
	s[0] = "aaa"
	fmt.Println("s", s, "s2", s2)

	good := []string{"A", "B", "C"}
	fmt.Println("good", good)
}
