package main

import (
	"fmt"
	"math"
)

// 基础知识：变量
func main() {
	var a = "hello"

	var b, c int = 1, 2

	var d = 3

	var e float32

	f := float64(e)

	g := a + " world"

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	const s string = "constant"
	const h = 520000
	const i = 1e10 / 100
	fmt.Println(s, math.Sin(h), math.Cos(i))
}
