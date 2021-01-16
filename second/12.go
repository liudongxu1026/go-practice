package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.通过 time 包获取当前时间
	nowTime := time.Now()
	fmt.Printf("%v\n", nowTime)
	// 2.将当前时间以下面两种格式打印出来
	// 2.1 2020-12-05 11:23:05
	// 2.1 2020-12-05 17:55:05
	time1 := nowTime.Format("2006-01-02 03:04:05")
	time2 := nowTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%v\n", time1)
	fmt.Printf("%v\n", time2)

	// 3.将当前时间转为时间戳 分别转换为 毫秒 纳秒
	second1 := nowTime.Unix()
	second2 := nowTime.UnixNano()
	fmt.Printf("%v\n", second1)
	fmt.Printf("%v\n", second2)
	// 4.将 1606621917 时间戳转换为时间
	var time3 int64 = 1606621917
	time3Obj := time.Unix(time3, 0)
	fmt.Printf("%v\n", time3Obj)

	// 5.将字符串时间分别转为 时间对象时间戳
	var timeStr string = "2020-11-29 11:51:57"
	timeStrObj, _ := time.ParseInLocation("2006-01-02 03:04:05", timeStr, time.Local)
	timeStrObjUnix := timeStrObj.Unix()
	fmt.Printf("%v\n", timeStrObj)
	fmt.Printf("%v\n", timeStrObjUnix)

	// 6.获取当前时间，给当前时间增加 1个小时 10分钟 20秒
	nowTime = nowTime.Add(time.Hour)
	fmt.Printf("%v\n", nowTime)
}
