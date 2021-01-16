package main

import "fmt"

func main() {
	// 1.实现选择排序
	arr1 := []int{5, 2, 6, 99, 3, 8}
	// for i := 0; i < len(arr1); i++ {
	// 	for j := 0; j < len(arr1); j++ {
	// 		if arr1[i] < arr1[j] {
	// 			arr1[i], arr1[j] = arr1[j], arr1[i]
	// 		}
	// 	}
	// }
	// fmt.Println(arr1)
	// 2.实现冒泡排序
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			fmt.Println(arr1[i], arr1[j])
			if arr1[i] > arr1[j] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}
	fmt.Println(arr1)
}
