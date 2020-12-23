package main

import "fmt"

// 1.定义一个函数，参数 x,y ，类型都是 int
// 	1.1 第一种
// func add(x int, y int) {

// }

// 	1.2 第二种
// func sub(x, y int) {

// }

// 2.定义一个函数 返回多个值 类型：int int
func addAndSubFn1(x, y int) (int, int) {
	return x + y, x - y
}

// 3.使用返回值命名法，返回多个值  类型：int int
func addAndSubFn2(x, y int) (add2 int, sub2 int) {
	add2 = x + y
	sub2 = x - y
	return
}

func main() {
	add1, sub1 := addAndSubFn1(1, 2)
	fmt.Printf("add1=%d sub1=%d\n", add1, sub1)
	add2, sub2 := addAndSubFn1(3, 4)
	fmt.Printf("add2=%d sub2=%d\n", add2, sub2)
}
