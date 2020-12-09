package main

import "fmt"

// func main() {
//  1.声明以下类型数组和切片 如下 可拓展
// 要求: 数组长度 10 位，切片长度 10 位 容量 30 位
// 打印 值 类型 长度 (切片要打印容量)
// [1,2,3] int
// ["1", "2", "3"] string
// [false, true, false] bool
// 	var arr1 = [10]int{}
// 	var arr2 = [10]float64{}
// 	var arr3 = [10]string{}
// 	var arr4 = [10]bool{}
// 	var arr5 = [10]byte{}
// 	var arr6 = [10]rune{}
// 	fmt.Printf("arr1 = %v 类型 = %T 长度 = %d\n", arr1, arr1, len(arr1))
// 	fmt.Printf("arr2 = %v 类型 = %T 长度 = %d\n", arr2, arr2, len(arr2))
// 	fmt.Printf("arr3 = %v 类型 = %T 长度 = %d\n", arr3, arr3, len(arr3))
// 	fmt.Printf("arr4 = %v 类型 = %T 长度 = %d\n", arr4, arr4, len(arr4))
// 	fmt.Printf("arr5 = %v 类型 = %T 长度 = %d\n", arr5, arr5, len(arr5))
// 	fmt.Printf("arr6 = %v 类型 = %T 长度 = %d\n", arr6, arr6, len(arr6))
// 	var slice1 = make([]int, 10, 30)
// 	var slice2 = make([]float64, 10, 30)
// 	var slice3 = make([]string, 10, 30)
// 	var slice4 = make([]bool, 10, 30)
// 	var slice5 = make([]byte, 10, 30)
// 	var slice6 = make([]rune, 10, 30)
// 	fmt.Printf("slice1 = %v 类型 = %T 长度 = %d 容量 = %d\n", slice1, slice1, len(slice1), cap(slice1))
// 	fmt.Printf("slice2 = %v 类型 = %T 长度 = %d 容量 = %d\n", slice2, slice2, len(slice2), cap(slice2))
// 	fmt.Printf("slice3 = %v 类型 = %T 长度 = %d 容量 = %d\n", slice3, slice3, len(slice3), cap(slice3))
// 	fmt.Printf("slice4 = %v 类型 = %T 长度 = %d 容量 = %d\n", slice4, slice4, len(slice4), cap(slice4))
// 	fmt.Printf("slice5 = %v 类型 = %T 长度 = %d 容量 = %d\n", slice5, slice5, len(slice5), cap(slice5))
// 	fmt.Printf("slice6 = %v 类型 = %T 长度 = %d 容量 = %d\n", slice6, slice6, len(slice6), cap(slice6))
// }

func main() {
	// 2.使用 for 遍历 10 以内数字
	// 0,1,2,3,4,5,6,7,8,9
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d\n", i)
	// }
	// 3.使用 for 遍历打印 99 表 (如下格式)
	// 1*1=1
	// 1*2=2 2*2=4
	// 1*3=3 2*3=6 3*3=9
	// for i := 1; i <= 9; i++ {
	// 	fmt.Printf("\n")
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%v * %v = %v  ", j, i, i*j)
	// 	}
	// }
	// 4.使用 for 打印 100 以内 奇数、偶数，存放到切片中
	// var oddSlice = make([]int, 0, 50)
	// var evenSlice = make([]int, 0, 50)
	// for i := 0; i <= 100; i++ {
	// 	if i%2 == 1 {
	// 		oddSlice = append(oddSlice, i)
	// 		fmt.Printf("奇数：%v ", i)
	// } else {
	// 		evenSlice = append(evenSlice, i)
	// 		fmt.Printf("偶数：%v ", i)
	// 	}
	// }
	// fmt.Printf("奇数：%v\n", oddSlice)
	// fmt.Printf("偶数：%v\n", evenSlice)
	// 5.label 循环 提示：配合 if break 使用
label:
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%v\n", i)
		for j := 0; j < 5; j++ {
			if j == 3 {
				break label //直接跳出两层循环
			}
			fmt.Printf("j=%v\n", j)
		}
	}
}
