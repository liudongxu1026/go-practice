package main

import "fmt"

// 1.声明一个公有结构体，声明结构体的方法并且调用
type Cat struct {
	Name  string
	Age   int
	Color string
}

// 两个方法 1个接受值 另一个接受指针类型的值
func (thisCat Cat) setName(name string) {
	thisCat.Name = name
}
func (thisCat *Cat) setName2(name string) {
	thisCat.Name = name
}

func main() {
	cat1 := Cat{
		Name:  "花花",
		Age:   10,
		Color: "白",
	}
	cat1.setName("花1")
	// setName {花花 10 白}
	fmt.Printf("setName %v\n", cat1)
	cat1.setName2("花2")
	// setName2 {花2 10 白}
	fmt.Printf("setName2 %v\n", cat1)
}
