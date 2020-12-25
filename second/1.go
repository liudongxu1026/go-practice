package main

import "fmt"

func main() {
	// 1.声明以下类型变量 使用 var 和 类型推导
	// int float string bool byte rune
	var int1 int64 = 1
	var float1 float64 = 1.2
	var str1 string = "abc"
	var bool1 bool = false
	var byte1 byte = 'a'
	var rune1 rune = '红'

	fmt.Printf("int1 %v %T\n", int1, int1)
	fmt.Printf("float1 %v %T\n", float1, float1)
	fmt.Printf("str1 %v %T\n", str1, str1)
	fmt.Printf("bool1 %v %T\n", bool1, bool1)
	fmt.Printf("byte1 %v %T\n", byte1, byte1)
	fmt.Printf("rune1 %v %T\n", rune1, rune1)

	int2 := 2
	float2 := 2.2
	str2 := "edf"
	bool2 := true
	byte2 := 'b'
	rune2 := '绿'
	fmt.Printf("int2 %v %T\n", int2, int2)
	fmt.Printf("float2 %v %T\n", float2, float2)
	fmt.Printf("str2 %v %T\n", str2, str2)
	fmt.Printf("bool2 %v %T\n", bool2, bool2)
	fmt.Printf("byte2 %v %T\n", byte2, byte2)
	fmt.Printf("rune2 %v %T\n", rune2, rune2)

	// 2.使用 Sprintf 拼接下方字符串
	// This is practice for golang
	myStr := fmt.Sprintf("This is %s for %s !", "practice", "golang")
	fmt.Printf("%v\n", myStr)
}
