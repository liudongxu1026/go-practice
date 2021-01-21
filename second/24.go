package main

import (
	"encoding/json"
	"fmt"
)

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
		Students: make([]Student, 0, 4),
	}
	// 2.1 将下列切片 nameArr 内容通过循环遍历插入到 Class.Students 字段中
	// 2.2 Student 的 gender 字段，通过循环 i 的取余决定
	nameArr := []string{"张三", "李四", "王五", "赵六"}
	sex := []string{"男", "女"}
	for index, oneName := range nameArr {
		enClass.Students = append(enClass.Students, Student{
			Name:   oneName,
			Age:    index * 10,
			Gender: sex[index%2],
		})
	}
	// 3.将这个实例化的结构体转为 json
	byte1, _ := json.Marshal(enClass)
	str1 := string(byte1)
	fmt.Printf("%v\n", str1)
	// 4.再把上方的json转为结构体
	enClass2 := MyClass{}
	json.Unmarshal([]byte(str1), &enClass2)
	fmt.Printf("%+v\n", enClass2)
}
