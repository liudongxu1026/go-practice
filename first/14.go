package main

import "fmt"

func main() {
	// 1.查看 a 变量的地址，取址
	var a = 10
	fmt.Printf("%p\n", &a) //0xc000012090

	// 2.声明一个新变量 p，将 a 的地址赋值给 p
	// 分别查看 变量地址 变量值 变量类型
	var p = &a
	// a地址：0xc000012090 a的值：10 a的类型：int
	fmt.Printf("a地址：%p a的值：%v a的类型：%T\n", &a, a, a)
	// p地址：0xc000006030 p的值：0xc000012090 p的类型：*int
	fmt.Printf("p地址：%p p的值：%v p的类型：%T\n", &p, p, p)

	// 3.声明一个变量 aValue ，值为 p 的指针
	var aValue = *p
	fmt.Printf("%v\n", aValue) //10
	// 通过对 p 取址修改值为 999
	*p = 999
	// 分别查看 a p aValue 值的变化
	fmt.Printf("a=%v\n", a)           //a=999
	fmt.Printf("p=%v\n", p)           //p=0xc000012090
	fmt.Printf("p=%v\n", *p)          //p=999
	fmt.Printf("aValue=%v\n", aValue) //aValue=10
}
