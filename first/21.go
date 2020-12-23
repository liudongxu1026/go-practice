package main

import "fmt"

// 1.声明两个结构体 a b
type Phone struct {
	name string
	info PhoneInfo
}

type PhoneInfo struct {
	money int
	color string
}

func main() {
	// 2.实例化嵌套结构体
	phone1 := Phone{
		name: "大哥大",
		info: PhoneInfo{
			money: 99999,
			color: "黑",
		},
	}
	fmt.Printf("%+v\n", phone1) //{name:大哥大 info:{money:99999 color:黑}}
}
