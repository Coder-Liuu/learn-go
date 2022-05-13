package main

import "fmt"

// 反射前夕
type Book struct {
}

// 定义两个接口
type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

// 使用Book的两个方法
func (this *Book) ReadBook() {
	fmt.Println("Read Book")
}
func (this *Book) WriteBook() {
	fmt.Println("Write Book")
}

func main() {
	// pair<type:Book, value:book>
	b := Book{}

	var r Reader
	// pair<type:Book, value:book>
	r = &b
	r.ReadBook()

	var w Writer
	// pair<type:Book, value:book>
	w = r.(Writer) // 强制类型转换
	w.WriteBook()
}
