package main

import (
	"fmt"
)

// 定义函数
func foo(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	var c = 100
	return c
}

func foo2() (int, int) {
	return 555, 666
}

func foo3() (r1 int, r2 int) {
	fmt.Println("\n----foo3----")
	r1 = 1000
	r2 = 1000
	return
}

// 变量类型后置
func add2(a *int) {
	*a += 2
}

func main() {
	c := foo("Hello", 100)
	fmt.Println("c = ", c)

	var ret1, ret2 = foo2()
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	ret1, ret2 = foo3()
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	a := 2
	add2(&a)
	fmt.Println(a)
}
