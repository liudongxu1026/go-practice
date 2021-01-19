package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string  `json:"name"`
	Age   float64 `json:"age"`
	Color string  `json:"color"`
}

func main() {
	// 1.使用Cat结构体，实例化结构体，将实例化的结构体转为 json (提示：使用json包)
	// 2.使用自定义 json 序列号的 key ，将实例化结构体转为 json
	oneCat := Cat{
		Name:  "小黑",
		Age:   3,
		Color: "black",
	}
	byte1, _ := json.Marshal(oneCat)
	json1 := string(byte1)
	fmt.Printf("%T %s\n", json1, json1)

	// 3.将字符串 json 转为实例化结构体
	str1 := `{"Name":"小黄","Age":1,"Color":"yellow"}`
	cat2 := Cat{}
	json.Unmarshal([]byte(str1), &cat2)
	fmt.Printf("%+v\n", cat2)
}
