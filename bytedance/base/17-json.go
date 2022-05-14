package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	var a = userInfo{"DWJ", 18, []string{"123", "234"}}

	// Marshal 转换为json格式
	buf, err := json.Marshal(a)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("buf:", buf)         // [123 34 78 97...]
	fmt.Println("buf:", string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	// Marshal + 排版
	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	// Unmarshal 转换为对象
	var b userInfo
	json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
}
