package main

import "fmt"

// 1.声明一个公有结构体、声明一个私有结构体
type User struct {
	name string
	age  int
	sex  string
}
type animal struct {
	name  string
	age   int
	color string
}

func main() {
	// 2.实例化上方声明的结构体
	// 2.1 请使用至少3种实例化结构体的方式（要包括普通赋值和指针赋值）
	user1 := User{
		name: "张三",
		age:  1,
		sex:  "男",
	}
	user2 := &User{ //指针赋值
		name: "李四",
		age:  2,
		sex:  "女",
	}
	user3 := User{}
	user3.name = "王五"
	user3.age = 3
	user3.sex = "女"
	fmt.Printf("%v\n", user1) //{张三 1 男}
	fmt.Printf("%v\n", user2) //&{李四 2 女}
	fmt.Printf("%v\n", user3) //{王五 3 女}

	// 3.声明一个使用类型默认值实例化结构体
	animal1 := animal{}
	fmt.Printf("%v\n", animal1) //{ 0 }

	// 4.请使用一段代码证明结构体是值类型
	animal2 := animal{
		name:  "小黄",
		age:   1,
		color: "yellow",
	}
	animal3 := animal2
	animal3.name = "small yellow"
	fmt.Printf("%v\n", animal2) // {小黄 1 yellow}
	fmt.Printf("%v\n", animal3) // {small yellow 1 yellow}
}
