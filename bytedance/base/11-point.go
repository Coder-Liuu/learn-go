package main

import "fmt"

func add2_(a int) {
	a += 2
}

func add2ptr(a *int) {
	*a += 2
}

func main() {
	var n = 5
	add2_(n)
	fmt.Println(n)

	add2ptr(&n)
	fmt.Println(n)
}
