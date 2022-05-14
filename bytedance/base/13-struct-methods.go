package main

import "fmt"

type Book struct {
	name string
}

func (this *Book) hello_ptr(word string) {
	fmt.Println(word, this.name)
}

func (this *Book) hello(word string) {
	fmt.Println(word, this.name)
}

func main() {
	var book = Book{"Book"}
	book.hello_ptr("Hello_Ptr")
	book.hello("Hello")
}
