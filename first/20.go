package main

import "fmt"

// 1.声明一个结构体
// 要求结构体种 有 stirng int 切片 map
type Cat struct {
	name     string
	age      int
	children []string
	info     map[string]string
}

func main() {
	// 2.实例化两个结构体
	// 第一个 切片和map要求是零值，不分配内存
	cat1 := Cat{
		name: "花花",
		age:  1,
	}
	fmt.Printf("%+v\n", cat1)     //{name:花花 age:1 children:[] info:map[]}
	fmt.Printf("%p\n", cat1.info) //0x0
	// 第二个 切片和map要求分配内存
	cat2 := Cat{
		name:     "喵喵",
		age:      2,
		children: make([]string, 3, 5),
		info:     make(map[string]string),
	}
	fmt.Printf("%+v\n", cat2)     //{name:喵喵 age:2 children:[  ] info:map[]}
	fmt.Printf("%p\n", cat2.info) // 0xc00006e360

}
