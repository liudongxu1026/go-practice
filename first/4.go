package main

// 提示：1 和 2 案例要使用 copy() 方法

func main() {
	// 1.复制切片 将切片A复制到切片B中
	// sliceA := []int{1, 2, 3, 4, 5}
	// sliceB := make([]int, 5)
	// copy(sliceB, sliceA)
	// fmt.Printf("sliceA=%v\n", sliceA) //[1 2 3 4 5]
	// fmt.Printf("sliceB=%v\n", sliceB) //[1 2 3 4 5]

	// 2.证明 数组和切片 的数据类型 是引用类型还是值类型？
	// sliceA[0] = 111
	// fmt.Printf("sliceA%v\n", sliceA)  //sliceA[111 2 3 4 5]
	// fmt.Printf("sliceB=%v\n", sliceB) //sliceB=[1 2 3 4 5]
	// 答2.1 切片是值类型，修改sliceA的第0个，sliceB没有被影响
	// arrA := [5]int{1, 2, 3, 4, 5}
	// arrB := arrA
	// fmt.Printf("arrA=%v\n", arrA) //[1 2 3 4 5]
	// fmt.Printf("arrB=%v\n", arrB) //[1 2 3 4 5]
	// arrA[0] = 111
	// fmt.Printf("arrA=%v\n", arrA) //arrA=[111 2 3 4 5]
	// fmt.Printf("arrB=%v\n", arrB) //arrB=[1 2 3 4 5]
	// 答2.2 数组是值类型，修改arrA的第0个，arrB没有被影响

	// 3.给切片扩容2倍
	// sliceC := make([]int, 5)
	// sliceC = append(sliceC, 1) // append会自动扩容
	// fmt.Printf("值=%v 长度=%d 容量=%d", sliceC, len(sliceC), cap(sliceC))

	// 4.声明1个10位的切片，取切片的前3位，后4位，中间6位
	// sliceD := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Printf("前3位=%v\n", sliceD[:3])
	// fmt.Printf("后4位=%v\n", sliceD[len(sliceD)-4:])
	// fmt.Printf("中间6位%v\n", sliceD[3:9])

	// 5.删除切片中索引为 2 的数值
	// sliceE := []int{0, 1, 2, 3, 4}

	// 6. ... 在切片前是干啥的？... 在切片后是干啥的？

}
