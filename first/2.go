package main

import (
	"fmt"
	"strconv"
)

// 提示：使用 strconv 包

func main() {
	// 1.将 int 类型转 string 类型
	var i int64 = 20
	str1 := strconv.FormatInt(i, 10)
	fmt.Printf("str1 = %s 类型 = %T\n", str1, str1)
	// 2.将 float 类型 转 string 类型
	var f float64 = 27.67654765
	fmt.Printf("f = %.2f 类型 = %T\n", f, f)
	str2 := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Printf("str2 = %s 类型 = %T\n", str2, str2)

	// 3.将 string 转 int 类型
	var str3 string = "1234"
	int1, _ := strconv.ParseInt(str3, 10, 64)
	fmt.Printf("int1 = %d 类型 = %T\n", int1, int1)

	// 4.将 string 转 float 类型
	var str4 string = "1234"
	float1, _ := strconv.ParseFloat(str4, 64)
	fmt.Printf("float1 = %.2f 类型 = %T\n", float1, float1)
}
