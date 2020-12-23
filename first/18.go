package main

import "fmt"

// 1.声明一个公有结构体，声明结构体的方法并且调用
// 两个方法 1个接受值 另一个接受指针类型的值
type Animal struct {
	name  string
	age   int
	color string
}

// 接收值
func (aniaml Animal) printIt() {
	fmt.Printf("%+v\n", aniaml)
}

// 接收指针
func (aniaml *Animal) editName(name string) {
	aniaml.name = name
}

func main() {
	cat := Animal{
		name:  "花花",
		age:   1,
		color: "白",
	}
	cat.printIt() //{name:花花 age:1 color:白}
	cat.editName("小白")
	cat.printIt() //{name:小白 age:1 color:白}
}
