package main

import "fmt"

// 1.声明一个自定义类型
type MyString string

// 2.给这个类型定义方法
func (str MyString) printIt() {
	fmt.Printf("%v\n", str)
}

// 3.调用
func main() {
	var str1 MyString = "哈哈"
	str1.printIt()
}
