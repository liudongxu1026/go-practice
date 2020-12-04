package main

import "fmt"

func main() {
	// 1.声明以下类型变量 使用 var 和 类型推导
	// int float string bool byte rune
	// var int1 int = 1
	// var float1 float64 = 0.1
	// var string1 string = "111"
	// var bool1 bool = true
	// var byte1 byte = 'a'
	// var rune1 rune = '哈'
	// fmt.Printf("int1：%v, 类型：%T\n", int1, int1)
	// fmt.Printf("float1：%v, 类型：%T\n", float1, float1)
	// fmt.Printf("string1：%v, 类型：%T\n", string1, string1)
	// fmt.Printf("bool1：%v, 类型：%T\n", bool1, bool1)
	// fmt.Printf("byte1：%v, 类型：%T\n", byte1, byte1)
	// fmt.Printf("rune1：%v, 类型：%T\n", rune1, rune1)
	// int2 := 2
	// float2 := 0.2
	// string2 := "222"
	// bool2 := false
	// byte2 := 'b'
	// rune2 := '嘿'
	// fmt.Printf("int2：%v, 类型：%T\n", int2, int2)
	// fmt.Printf("float2：%v, 类型：%T\n", float2, float2)
	// fmt.Printf("string2：%v, 类型：%T\n", string2, string2)
	// fmt.Printf("bool2：%v, 类型：%T\n", bool2, bool2)
	// fmt.Printf("byte2：%v, 类型：%T\n", byte2, byte2)
	// fmt.Printf("rune2：%v, 类型：%T\n", rune2, rune2)
	// 2.使用 Sprintf 拼接下方字符串
	// This is practice for golang
	var str1 string = "This is practice for"
	var str2 string = "golang"
	var str3 string = fmt.Sprintf("%s %s 吼吼吼", str1, str2)
	fmt.Printf("%s", str3)
}
