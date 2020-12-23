package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.定时器 提示：使用NewTicker，它会返回值中有一个 C ，循环它
	// 使用定时器 5 次当前时间
	// ticker := time.NewTicker(time.Second)
	// num := 0
	// for time := range ticker.C {
	// 	if num == 5 {
	// 		ticker.Stop()
	// 		break
	// 	}
	// 	num++
	// 	fmt.Printf("%v\n", time)
	// }

	// 2.使用 死循环结合Sleep方法 实现定时器
	for {
		time.Sleep(time.Second)
		fmt.Print("打印1次\n")
	}
}
