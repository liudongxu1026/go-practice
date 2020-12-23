package main

import "fmt"

// {
// 1.封装 1 个函数，递归计算 参数(例：100) 之内的合，返回合
// 接受 1 个 int 类型，返回一个 int 类型
// func add(a int) int {
// 	if a > 1 {
// 		return add(a-1) + a
// 	} else {
// 		return 1
// 	}
// }

// func main() {
// 	result := add(100)
// 	fmt.Printf("100以内之合: %v", result)
// }
// }

// 2.两个函数，均返回一个闭包函数，函数返回一个 int 类型的值
// 每个函数中有一个 x := 10
// 2.1 第一个闭包 不接受参数 返回 x+1
func addFn1() func() int {
	x := 10
	return func() int {
		return x + 1
	}
}

// 2.2 第二个闭包 接受参数 返回 x+参数
func addFn2() func(x int) int {
	return func(x int) int {
		return x + 1
	}
}

// 2.3 调用 3 次，对比结果
func main() {
	fn1 := addFn1()
	fmt.Printf("fn1: %v\n", fn1()) // 11
	fmt.Printf("fn1: %v\n", fn1()) // 11
	fmt.Printf("fn1: %v\n", fn1()) // 11
	fn2 := addFn2()
	fmt.Printf("fn2: %v\n", fn2(20)) // 21
	fmt.Printf("fn2: %v\n", fn2(30)) // 31
	fmt.Printf("fn2: %v\n", fn2(40)) // 41
}
