package main

import (
	"errors"
	"fmt"
)

type Cat struct {
	name string
}

func findCat(cats []Cat, name string) (*Cat, error) {
	for _, cat := range cats {
		if cat.name == name {
			return &cat, nil
		}
	}
	return nil, errors.New("Not Found")
}

func main() {
	var cats = []Cat{{"A"}, {"B"}, {"C"}}

	if cat, err := findCat(cats, "A"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("cat:", cat.name)
	}

	if cat, err := findCat(cats, "D"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("cat:", cat.name)
	}
}
