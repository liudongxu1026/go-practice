package main

import "fmt"

func main() {
	// 1.声明以下类型数组和切片 如下 可拓展
	// 要求: 数组长度 10 位，切片长度 10 位 容量 30 位
	// 打印 值 类型 长度 (切片要打印容量)
	// [1,2,3] int
	// ["1", "2", "3"] string
	// [false, true, false] bool
	var arr1 = [10]int{}
	var arr2 = [10]string{}
	var arr3 = [10]bool{}
	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", arr3)
	var slice1 = make([]int, 10, 30)
	var slice2 = make([]string, 10, 30)
	var slice3 = make([]bool, 10, 30)
	fmt.Printf("%v\n", slice1)
	fmt.Printf("%v\n", slice2)
	fmt.Printf("%v\n", slice3)

	// 2.使用 for 遍历 10 以内数字
	// 0,1,2,3,4,5,6,7,8,9
	for i := 0; i < 10; i++ {
		fmt.Printf("%v,", i)
	}
	fmt.Printf("\n")
	// 3.使用 for 遍历打印 99 表 (如下格式)
	// 1*1=1
	// 1*2=2 2*2=4
	// 1*3=3 2*3=6 3*3=9
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", j, i, j*i)
		}
		fmt.Printf("\n")
	}
	// 4.使用 for 打印 100 以内 奇数、偶数，存放到切片中
	var oddSlice = make([]int, 0, 50)
	var evenSlice = make([]int, 0, 50)
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenSlice = append(evenSlice, i)
		} else {
			oddSlice = append(oddSlice, i)
		}
	}
	fmt.Printf("oddSlice = %v\n", oddSlice)
	fmt.Printf("evenSlice = %v\n", evenSlice)
	// 5.label 循环 提示：配合 if break 使用
thisFor:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%v %v\n", i, j)
			if j == 3 {
				break thisFor
			}
		}
	}
}
