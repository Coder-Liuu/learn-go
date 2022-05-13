package main

import "fmt"

func main() {
	a := 2

	// switch中不需要在每个case中写break
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// if中不需要写()
	if a == 2 {
		fmt.Println("if a == 2")
	} else {
		fmt.Println("if a != 2")
	}
}
