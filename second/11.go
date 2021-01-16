package main

import "fmt"

// {
// 1.声明两个函数 查看 retur 和 defer 执行的时机
// 1.1 第一个函数中 声明一个变量 a， defer一个函数，函数中执行 a++，最后由外层函数返回 a
// func fn1() int {
// 	var a int
// 	defer func() {
// 		a++
// 	}()
// 	return a
// }

// 1.2 第二个函数返回一个变量是a的值，defer一个内部函数，函数中执行 a++，最后由外层函数返回 a
// func fn2() (a int) {
// 	defer func() {
// 		a++
// 	}()
// 	return a
// }

// 1.3 查看上方执行不同
// func main() {
// 	fmt.Printf("fn1 %v\n", fn1())
// 	fmt.Printf("fn2 %v\n", fn2())
// }
// }

// 2.使用defer panic抛出异常 提示：要使用 recover() 参考 22.go
func sendError() {
	defer func() {
		err := recover()
		fmt.Printf("%v++++\n", err)
	}()
	panic("报错了111")
}
func main() {
	fmt.Println("开始")
	sendError()
	fmt.Println("结束")
}
