package main

import "fmt"

// 1.定义一个函数，参数 x,y ，类型都是 int
// 	1.1 第一种
func fn1(x, y int) {}

// 	1.2 第二种
func fn2(x int, y int) {}

// 2.定义一个函数 返回多个值 类型：int int
func fn3() (int, int) {
	return 1, 2
}

// 3.使用返回值命名法，返回多个值  类型：int int
func fn4() (c int, d int) {
	c = 3
	d = 4
	return
}
func main() {
	var a, b = fn3()
	fmt.Printf("%v %v\n", a, b)
	var c, d = fn4()
	fmt.Printf("%v %v\n", c, d)
}
