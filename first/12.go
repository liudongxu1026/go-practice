package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.通过 time 包获取当前时间
	nowTime := time.Now()
	// fmt.Printf("%v\n", nowTime)
	// 2.将当前时间以下面两种格式打印出来
	// 2.1 2020-12-05 11:23:05
	// 2.1 2020-12-05 17:55:05
	// 2016 03 04 05
	time1 := nowTime.Format("2006-01-02 03:04:05")
	time2 := nowTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%v\n", time1) //2020-12-18 04:51:46
	fmt.Printf("%v\n", time2) //2020-12-18 16:51:46

	// 3.将当前时间转为时间戳 分别转换为 毫秒 纳秒
	time3 := nowTime.Unix()     //毫秒
	fmt.Printf("%v\n", time3)   //1608525373
	time4 := nowTime.UnixNano() //纳秒
	fmt.Printf("%v\n", time4)   //1608525420254568800

	// 4.将 1606621917 时间戳转换为时间
	var time5 int64 = 1606621917
	time5Obj := time.Unix(time5, 0)
	fmt.Printf("%v\n", time5Obj)

	// 5.将字符串时间分别转为 时间对象和时间戳
	var timeStr string = "2020-11-29 11:51:57"
	time6Obj, _ := time.ParseInLocation("2006-01-02 03:04:05", timeStr, time.Local)
	fmt.Printf("%v\n", time6Obj)

	// 6.获取当前时间，给当前时间增加 1个小时 10分钟 20秒
	nowTime = nowTime.Add(time.Hour)
	fmt.Printf("%v\n", nowTime)

}
