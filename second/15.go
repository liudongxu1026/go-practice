package main

import "fmt"

// 1.声明两个函数 一个变量 a = 5
// 1.1 函数A 接受 a 变量 int 类型，对参数重新赋值
func fn1(a int) {
	a = 10
}

// 1.2 函数B 接受 a 变量的 int 指针类型，对参数重新赋值
func fn2(a *int) {
	*a = 20
}

func main() {
	// 1.3 查看外部变量 a 的变化
	a := 5
	fn1(a)
	fn2(&a)
	fmt.Printf("%v\n", a) //20
}
