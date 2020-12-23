package main

// {
// (不懂就看 it-go 16.go)
// 1.使用 type 定义 calc 方法类型
// type calc func(int, int) int

// 1.1 定义 1 个函数为 calc 类型
// func add(x int, y int) int {
// 	return x + y
// }

// 1.2 定义 1 个函数属于 main
// func sub(x, y int) int {
// 	return x - y
// }

// 1.3 打印查看上方 2 个函数的不同
// func main() {
// 	var addFn calc = add
// 	fmt.Printf("%T\n", addFn) //main.calc
// 	fmt.Printf("%T\n", sub)   //func(int, int) int
// }
// }

// {
// (不懂就看 it-go 17.go)
// 2.使用 type 定义 calcType 个方法类型 接受参数 int int，返回参数 int
// type calcType func(int, int) int

// 2.1 定义 calc 函数 1参 2参接受 int 3参接受 calcType 类型，返回值为 int
// func calc(a int, b int, fn calcType) int {
// 	return fn(a, b)
// }

// 2.2 定义 add 函数，参数类型与 calcType 类型相同，返回值为 int
// func add(a, b int) int {
// 	return a + b
// }

// func main() {
// 2.3 调用 calc 方法，返回 第 3 个参数的返回值
// 	var result = calc(3, 4, add)
// 	fmt.Printf("%v", result)
// }
// }
