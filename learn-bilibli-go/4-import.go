package main

import (
	// 但是不使用
	_ "learn-bilibli-go/3-init/lib1"
	// 导入起别名
	lib2 "learn-bilibli-go/3-init/lib2"
	// . "learn-bilibli-go/3-init/lib2"
)

func main() {
	lib2.Test()
}
