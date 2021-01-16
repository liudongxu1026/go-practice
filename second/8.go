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
	var allKey = make([]string, 0, 4)
	for key, _ := range map1 {
		allKey = append(allKey, key)
	}
	sort.Strings(allKey)
	for _, key := range allKey {
		fmt.Printf("%v %v\n", key, map1[key])
	}
}
