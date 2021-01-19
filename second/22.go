package main

import "fmt"

// 1.声明两个结构体 Animal Dog (提示：与嵌套结构体一样)
type Animal struct {
	Type string
}

func (thisAniaml Animal) printAniamlName() {
	fmt.Printf("%v\n", thisAniaml.Type)
}

func (thisDog Dog) printDogName() {
	fmt.Printf("%v\n", thisDog.Name)
}

type Dog struct {
	Name string
	Age  int
	Animal
}

func main() {
	// 2.要求实例化 Dog 结构体，并嵌套实例化 Animal，继承 Animal 的方法
	dog1 := Dog{
		Name: "小黄",
		Age:  10,
		Animal: Animal{
			Type: "狗子",
		},
	}
	dog1.printAniamlName()
	dog1.printDogName()
}
