package main

import "fmt"

// 1.封装 1 个函数，递归计算 参数(例：100) 之内的合，返回合
// 接受 1 个 int 类型，返回一个 int 类型
func calc(total int) int {
	if total > 1 {
		return calc(total-1) + total
	} else {
		return 1
	}
}

// 2.两个函数，均返回一个闭包函数，函数返回一个 int 类型的值
// 每个函数中有一个 x := 10
// 2.1 第一个闭包 不接受参数 返回 x+1
func fn1() func() int {
	x := 10
	return func() int {
		return x + 1
	}
}

// 2.2 第二个闭包 接受参数x 返回 x+1
func fn2() func(int) int {
	return func(x int) int {
		return x + 1
	}
}

// 2.3 调用 3 次，对比结果
func main() {
	var res = calc(100)
	fmt.Printf("res = %v\n", res)
	var myFn1 = fn1()
	fmt.Printf("fn1 = %v\n", myFn1())
	fmt.Printf("fn1 = %v\n", myFn1())
	fmt.Printf("fn1 = %v\n", myFn1())
	var myFn2 = fn2()
	fmt.Printf("fn1 = %v\n", myFn2(10))
	fmt.Printf("fn1 = %v\n", myFn2(20))
	fmt.Printf("fn1 = %v\n", myFn2(30))
}
