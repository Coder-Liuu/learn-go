package main

import (
	"fmt"
	"time"
)

func main() {
	go func(a int, b int) {
		fmt.Printf("a = %d, b = %d", a, b)
	}(10, 20)

	time.Sleep(1000000)
}
