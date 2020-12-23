package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1.声明一个 map 类型，键为 string 类型，值为 string 类型
	// cat1 := map[string]string{
	// 	"name":  "猫猫",
	// 	"age":   "1",
	// 	"color": "black",
	// }
	// fmt.Printf("%#v\n", cat1)

	// 2.判断 id 这个键是否存在(在上一个map类型)？打印出 bool 类型的值
	// idVal, hasId := cat1["id"]
	// nameVal, hasName := cat1["name"]
	// fmt.Printf("id %v %v\n", idVal, hasId)
	// fmt.Printf("name %v %v\n", nameVal, hasName)

	// 3.在上面 map 中删除 1 个字段
	// delete(cat1, "age")
	// fmt.Printf("%#v\n", cat1)

	// 4.声明 1 个切片，值为 map 类型，键为 string 类型，值为 string 类型
	// sliceMap := make([]map[string]string, 3, 5)
	// fmt.Printf("%v\n", sliceMap)

	// 5.查看单词出现频率 提示:要使用 strings 包
	// 5.1 小拓展：封装一个方法，通过传参决定是否区分大小写
	count1 := getStrCount(true)
	count2 := getStrCount(false)
	fmt.Printf("%v\n", count1)
	fmt.Printf("%v\n", count2)
}

// 方法默认是区分大小写
func getStrCount(isDif bool) map[string]int {
	var words string = "How do you do How Do You Do how do you do"
	var wordSplit = strings.Split(words, " ")
	var count = make(map[string]int)
	for _, val := range wordSplit {
		thisWord := val
		// 不区分大小写
		if !isDif {
			thisWord = strings.ToLower(val)
		}
		_, hasNowKey := count[thisWord]
		if hasNowKey {
			count[thisWord]++
		} else {
			count[thisWord] = 1
		}
	}
	return count
}
