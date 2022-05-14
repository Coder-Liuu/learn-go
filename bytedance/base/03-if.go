package main

import "fmt"

//基础知识：if语句
func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 8; num < 0 {
		fmt.Println("8 < 0")
	} else if num == 0 {
		fmt.Println("8 == 0")
	} else {
		fmt.Println("8 > 0")
	}
}
