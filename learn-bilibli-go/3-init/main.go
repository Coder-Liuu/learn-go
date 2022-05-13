package main

import (
	"learn-bilibli-go/3-init/lib1"
	"learn-bilibli-go/3-init/lib2"
)

func main() {
	// Golang 文件内，首字母大写的函数、常量等
	// 可以被其它模块访问，首字母小写的只能在当前文件中可以使用。
	lib1.Test()
	lib2.Test()
}
