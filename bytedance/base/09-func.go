package main

import "fmt"

func add1(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]int, key string) (int, bool) {
	val, ok := m[key]
	return val, ok
}

func main() {
	res1 := add1(1, 2)
	res2 := add2(3, 4)
	fmt.Println("res1", res1, "res2", res2)

	var m = map[string]int{"one": 1, "two": 2}
	fmt.Println(exists(m, "unknow"))
}
