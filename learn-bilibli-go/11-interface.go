package main

import "fmt"

/*
步骤:
    1. 有一个父类接口
    2. 实现了所有方法
    3. 指向子类
*/

// 动物类接口
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Cat 对象
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

// Dog 对象
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animal Animal) {
	fmt.Println("Type = ", animal.GetType())
	fmt.Println("Color = ", animal.GetColor())
	animal.Sleep()
	fmt.Println()
}

func main() {
	// 方式1
	var animal Animal // 定义一个父类指针，指向子类的地址
	fmt.Println("\n方式1")
	animal = &Cat{"Black"}
	animal.Sleep()

	animal = &Dog{"Green"}
	animal.Sleep()

	// 方式2
	fmt.Println("\n方式2")
	a1 := Cat{"Black"}
	a2 := Dog{"Green"}
	ShowAnimal(&a1)
	ShowAnimal(&a2)
}
