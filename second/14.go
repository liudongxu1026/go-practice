package main

import "fmt"

func main() {
	// 1.查看 a 变量的地址，取址
	var a = 10
	fmt.Printf("%p\n", &a)
	// 2.声明一个新变量 p，将 a 的地址赋值给 p
	// 分别查看 变量地址 变量值 变量类型
	var p = &a
	// fmt.Printf("a=%v 地址=%p 类型=%T\n", a, &a, a)
	// fmt.Printf("p=%v 地址=%p 类型=%T\n", p, &p, p)
	// 3.声明一个变量 aValue ，值为 p 的指针
	// 通过对 p 取址修改值为 999
	// 分别查看 a p aValue 值的变化
	var aValue = &p
	*p = 999
	fmt.Printf("a=%v 地址=%p 类型=%T\n", a, &a, a)
	fmt.Printf("p=%v 地址=%p 类型=%T\n", p, &p, p)
	fmt.Printf("aValue=%v 地址=%p 类型=%T\n", aValue, &aValue, aValue)

}
