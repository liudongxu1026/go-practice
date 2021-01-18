package main

import "fmt"

func main() {
	// 1.使用 make 声明一个键值为 string 的 map 类型
	map1 := make(map[string]string)
	fmt.Printf("map1 = %v\n", map1)
	// 2.使用 make 声明一个切片，类型为 int ，长度为 10，容量为 15
	slice1 := make([]int, 10, 15)
	fmt.Printf("slice1 = %v 长度 = %d 容量 = %d\n", slice1, len(slice1), cap(slice1))
	// 3.使用 new 分配一个存 int 类型数据的地址，变量名为 a
	a := new(int)
	fmt.Printf("%v %T %p\n", *a, a, a)
	// 通过地址取值操作 将 a 值修改为 10
	*a = 10
	// 查看 a 变量自己的地址
	fmt.Printf("%p\n", &a)
	// 查看 a 值保存的地址
	fmt.Printf("%v %p\n", *a, a)
}
