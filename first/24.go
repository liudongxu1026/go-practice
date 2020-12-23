package main

import (
	"encoding/json"
	"fmt"
)

// 1.声明两个结构体 Class Student
type Class struct {
	Name     string
	Students []Student
}

type Student struct {
	Name   string
	Age    int
	gender string
}

func main() {
	// 2.实例化 Class 结构体
	mathClass := Class{
		Name:     "数学班级",
		Students: make([]Student, 0, 4),
	}
	// 2.1 将下列切片 nameArr 内容通过循环遍历插入到 Class.Students 字段中
	nameArr := []string{"张三", "李四", "王五", "赵六"}
	// 2.2 Studnet 的 gender 字段，通过循环 i 的取余决定
	sex := []string{"男", "女"}
	for index, oneName := range nameArr {
		oneStu := Student{
			Name:   oneName,
			Age:    (index + 1) * 2,
			gender: sex[index%2],
		}
		mathClass.Students = append(mathClass.Students, oneStu)
	}
	// fmt.Printf("类型：%T, 值：%+v\n", mathClass, mathClass)

	// 3.将这个实例化的结构体转为 json
	mathClassByte, _ := json.Marshal(mathClass)
	mathClassjson := string(mathClassByte)
	fmt.Printf("类型：%T, 值：%+v\n", mathClassjson, mathClassjson)
	// 4.再把上方的json转为结构体
	mathClassTwo := Class{}
	json.Unmarshal([]byte(mathClassjson), &mathClassTwo)
	fmt.Printf("%T %+v\n", mathClassTwo, mathClassTwo)
	fmt.Printf("%T %+v\n", mathClassTwo.Students, mathClassTwo.Students)
}
