package main

import (
	"fmt"
	"time"
)

// func main() {
// 1.定时器 提示：使用NewTicker，它会返回值中有一个 C ，循环它
// 使用定时器 5 次当前时间
// 	time := time.NewTicker(time.Second)
// 	num := 0
// 	for one := range time.C {
// 		if num == 5 {
// 			break
// 		}
// 		fmt.Printf("%v\n", one)
// 		num++
// 	}
// }

// 2.使用 死循环结合Sleep方法 实现定时器
func main() {
	for {
		fmt.Println("1")
		time.Sleep(time.Second)
	}
}
