package main

import "fmt"

// 1.声明两个结构体 Class Student
type MyClass struct {
	Subject  string
	Students []Student
}
type Student struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	// 2.实例化 Class 结构体
	enClass := MyClass{
		Subject:  "英语",
		Students: make([]Student, 4),
	}
	// fmt.Printf("%+v\n", enClass)
	// 2.1 将下列切片 nameArr 内容通过循环遍历插入到 Class.Students 字段中
	// nameArr := []string{"张三", "李四", "王五", "赵六"}
	for index, oneStu := range enClass.Students {
		fmt.Printf("%v %+v\n", index, oneStu)
	}
	// 2.2 Student 的 gender 字段，通过循环 i 的取余决定
	// sex := []string{"男", "女"}

	// 3.将这个实例化的结构体转为 json

	// 4.再把上方的json转为结构体
}
