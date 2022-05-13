package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 4, []string{"zhang3", "li4", "wang5"}}

	// 编码过程 struct -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal err")
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码过程 json -> struct
	var myMovie Movie
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal err")
		return
	}
	fmt.Printf("struct = %v\n", myMovie)
}
