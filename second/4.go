// 提示：1 和 2 案例要使用 copy() 方法
package main

import "fmt"

func main() {
	// 1.复制切片 将切片A复制到切片B中
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := make([]int, 5, 5)
	copy(sliceB, sliceA)
	fmt.Printf("sliceA %v\n", sliceA)
	fmt.Printf(" sliceB%v\n", sliceB)
	// 2.证明 数组和切片 的数据类型 是引用类型还是值类型？
	// 2.1 数组
	arr1 := [3]string{"a", "b", "c"}
	arr2 := arr1
	arr1[0] = "aaa"
	fmt.Printf("arr1 %v\n", arr1) //[aaa b c]
	fmt.Printf("arr2 %v\n", arr2) //[a b c]
	// 2.2 切片：引用类型
	sliceC := []int{1, 2, 3, 4, 5}
	sliceD := sliceC
	sliceC[0] = 111
	fmt.Printf("sliceC %v\n", sliceC) //[111 2 3 4 5]
	fmt.Printf("sliceD %v\n", sliceD) //[111 2 3 4 5]
	// 3.给切片扩容2倍
	sliceE := []int{1, 2, 3, 4, 5}
	// sliceE [1 2 3 4 5], 容量 5
	fmt.Printf("sliceE %v, 容量 %d\n", sliceE, cap(sliceE))
	sliceE = append(sliceE, 6)
	// sliceE [1 2 3 4 5 6], 容量 10
	fmt.Printf("sliceE %v, 容量 %d\n", sliceE, cap(sliceE))
	// 4.声明1个10位的切片，取切片的前3位，后4位，中间6位
	sliceF := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", sliceF)
	fmt.Printf("前3位 %v\n", sliceF[:3])
	fmt.Printf("后4位 %v\n", sliceF[6:])
	fmt.Printf("中间6位 %v\n", sliceF[2:8])
	// 5.删除切片中索引为 2 的数值
	sliceF = append(sliceF[:2], sliceF[3:]...)
	fmt.Printf("%v\n", sliceF)
	// 6. 使用 ... 合并和解散切片
	// 6.1 声明一个切片
	// 6.2 声明一个函数，传参时将切片打散，接收参数时将参数合并为切片
	sliceG := []string{"张三", "李四", "王五"}
	printName(sliceG...)
}

func printName(nameSlice ...string) {
	fmt.Printf("nameSlice %v\n", nameSlice)
}
