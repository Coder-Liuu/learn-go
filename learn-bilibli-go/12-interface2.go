package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("\nThis is", arg)

	// 返回values, ok  类型断言
	_, ok := arg.(int)
	if ok {
		fmt.Println("I am int")
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"zhang3"}
	myFunc(book)

	var i int = 1
	myFunc(i)

	var f32 float32 = 1.23
	myFunc(f32)
}
