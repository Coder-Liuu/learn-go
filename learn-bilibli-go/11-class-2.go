package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Walk() {
	fmt.Println("Human Walk...")
}

func (this *Human) Show() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
}

// 继承Human对象
type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Show() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhangsan", "male"}
	h.Show()

	fmt.Println("\n")
	s := SuperMan{Human{"li3", "male"}, 88}
	s.Show()
}
