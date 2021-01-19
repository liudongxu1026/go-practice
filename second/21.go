package main

import "fmt"

// 1.声明两个结构体 a b
type Father struct {
	Name  string
	Age   int
	Child Son
}

type Son struct {
	Name string
	Age  int
}

func main() {
	// 2.实例化嵌套结构体
	oneFather := Father{
		Name: "爸爸",
		Age:  40,
		Child: Son{
			Name: "儿子",
			Age:  10,
		},
	}
	fmt.Printf("%+v\n", oneFather)
}
