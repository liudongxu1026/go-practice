package main

import "fmt"

// 1.声明一个自定义类型
type myInt int

// 2.给这个类型定义方法
func (this myInt) printThis() {
	fmt.Printf("%v\n", this)
}
func main() {
	// 3.调用
	var a myInt = 10
	a.printThis() //10
}
