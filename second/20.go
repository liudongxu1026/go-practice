package main

import "fmt"

// 1.声明一个结构体
// 要求结构体中 有 stirng int 切片 map
type User struct {
	Name    string
	Hobbies []string
	Info    map[string]string
}

func main() {
	// 2.实例化两个结构体
	// 第一个 切片和map要求是零值，不分配内存
	user1 := User{
		Name: "张三",
	}
	fmt.Printf("%p %p\n", user1.Hobbies, user1.Info)
	// 第二个 切片和map要求分配内存
	user2 := User{
		Name:    "张三",
		Hobbies: make([]string, 0),
		Info:    make(map[string]string),
	}
	fmt.Printf("%p %p\n", user2.Hobbies, user2.Info)

}
