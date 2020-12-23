package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1.将下面 map ，按照 key 的 abcd... 顺序打印输出
	map1 := map[string]string{
		"username": "张三",
		"age":      "12",
		"sex":      "男",
		"height":   "150",
	}
	// 1.1 将 map 转成切片
	keySlice := make([]string, 0, 4)
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	// 1.2 将 key 排序
	sort.Strings(keySlice)
	// 1.2 输出打印
	fmt.Printf("%v\n", keySlice)
	for _, key := range keySlice {
		fmt.Printf("key=%v val=%v\n", key, map1[key])
	}
}
