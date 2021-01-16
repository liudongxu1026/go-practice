package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 1.声明一个 map 类型，键为 string 类型，值为 string 类型
	map1 := map[string]string{
		"name": "李四",
		"age":  "10",
		"id":   "001",
		"sex":  "女",
	}
	fmt.Printf("%v\n", map1)
	// 2.判断 id 这个键是否存在(在上一个map类型)？打印出 bool 类型的值
	val, hasId := map1["id"]
	fmt.Printf("%v %v %T\n", val, hasId, hasId)
	// 3.在上面 map 中删除 1 个字段
	delete(map1, "age")
	fmt.Printf("%v\n", map1)

	// 4.声明 1 个切片，值为 map 类型，键为 string 类型，值为 string 类型
	slice1 := make([]map[string]string, 0, 3)
	fmt.Printf("%v\n", slice1)

	// 5.查看单词出现频率 提示:要使用 strings 包
	// 5.1 小拓展：封装一个方法，通过传参决定是否区分大小写
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	var caseMap1 = wordNum(false) //不区分
	var caseMap2 = wordNum(true)  //区分
	fmt.Printf("%v\n", caseMap1)
	fmt.Printf("%v\n", caseMap2)
}

// isCase false为区分 true不区分
func wordNum(isCase bool) map[string]int {
	var words = "How do you do How Do You Do how do you do"
	var result = make(map[string]int)
	var wordsArr = strings.Split(words, " ")
	for _, key := range wordsArr {
		thisKey := key
		if !isCase {
			thisKey = strings.ToLower(key)
		}
		_, isHasKey := result[thisKey]
		if isHasKey {
			result[thisKey]++
		} else {
			result[thisKey] = 1
		}
	}
	return result
}
