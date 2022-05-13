package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Name string `info:"名字"`
	Sex  string `info:"性别"`
}

func reflectTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println("info = ", tagString)
	}
}

func main() {
	book := Book{"zhang3", "male"}
	reflectTag(&book)
}
