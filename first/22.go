package main

import "fmt"

// 1.声明两个结构体 Animal Dog (提示：与嵌套结构体一样)
type Animal struct {
	name string
}

func (a Animal) printAniaml() {
	fmt.Printf("%+v\n", a)
}
func (d Dog) dogFn() {
	fmt.Printf("%+v\n", d)
}

type Dog struct {
	age    int
	Animal //继承
}

func main() {
	// 2.要求实例化 Dog 结构体，并嵌套实例化 Animal，继承 Animal 的方法
	// 嵌套继承时要实现普通实例化
	dog1 := Dog{
		age: 1,
		Animal: Animal{
			name: "旺旺",
		},
	}
	dog1.dogFn()
	dog1.printAniaml()
}
