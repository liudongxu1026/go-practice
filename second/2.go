// 提示：使用 strconv 包
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1.将 int 类型转 string 类型
	var i int64 = 20
	str1 := strconv.FormatInt(i, 10)
	fmt.Printf("%v %T\n", str1, str1)
	// 2.将 float 类型 转 string 类型
	var f float32 = 27.67654765
	float1 := strconv.FormatFloat(float64(f), 'f', -1, 64)
	fmt.Printf("%v %T\n", float1, float1)
	// 3.将 string 转 int 类型
	var str string = "1234"
	int1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%v %T\n", int1, int1)

	// 4.将 string 转 float 类型
	var str3 string = "1234.33"
	float2, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("%v %T\n", float2, float2)

}
