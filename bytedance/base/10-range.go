package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4}
	for val, idx := range nums {
		fmt.Println(val, idx)
	}

	var map1 = map[string]int{"one": 1, "two": 2}
	for key, val := range map1 {
		fmt.Println(key, val)
	}

	for key := range map1 {
		fmt.Println("key", key)
	}
}
