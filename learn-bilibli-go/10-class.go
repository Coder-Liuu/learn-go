package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

// 有星为引用，无星为值传递
// 带指针可以对结构体进行修改，不带指针不可以对结构体修改
// 带指针可以避免对大结构体拷贝的开销
func (this *Hero) Show() {
	fmt.Println("=======")
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

// 大写为共有，小写为私有
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(NewName string) {
	this.Name = NewName
}

func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()
	hero.SetName("li4")
	hero.Show()
}
