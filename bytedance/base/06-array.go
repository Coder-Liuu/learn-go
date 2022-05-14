package main

import "fmt"

func main() {
	var a [5]int

	var b = [5]int{1, 2, 3, 4, 5}

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}

	fmt.Println("a", a, "\nb", b, "\nc", c)
}
