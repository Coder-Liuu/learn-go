package main

import (
	"fmt"
	"reflect"
)

// ===== 简单类型 =====
func reflectNum(arg interface{}) {
	fmt.Println("type = ", reflect.TypeOf(arg))
	fmt.Println("values = ", reflect.ValueOf(arg))
	fmt.Println()
}

// ===== 复杂类型 =====
type User struct {
	Id   int
	Name string
	Age  float32
}

func (this User) Call() {
	fmt.Println("User Call...")
}

func reflectUser(arg interface{}) {
	// 获取type
	inputType := reflect.TypeOf(arg)
	// 获取value
	inputValue := reflect.ValueOf(arg)

	fmt.Println("type = ", inputType) // floag
	fmt.Println("values = ", inputValue)
	// 获取所有字段
	fmt.Println("\n获取所有字段")
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// 获取所有方法
	fmt.Println("\n获取所有方法")
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	// 反射:把一个变量的原型照出来
	var num1 float64 = 12.234
	var num2 int32 = 12
	reflectNum(num1)
	reflectNum(num2)

	// 反射复杂类型
	var user User = User{2, "zhang3", 20}
	reflectUser(user)
}
