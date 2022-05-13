package main

import "fmt"
import "errors"

// 错误处理
func add(a int) (int, error) {
	if a == 2 {
		return 2, nil
	}
	return 3, errors.New("Not Found")

}

func main() {
	fmt.Println(add(2))
	fmt.Println(add(3))
	fmt.Println("End")
}
