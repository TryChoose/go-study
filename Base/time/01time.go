package main

import (
	"fmt"
	"time"
)

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func f1() {
	////获取当前时间
	//now := time.Now()
	//fmt.Println(now)
	////年月日 时分秒
	//fmt.Println(now.Year())
	//fmt.Println(now.Month())
	//fmt.Println(now.Date())
	//fmt.Println(now.Hour())
	//fmt.Println(now.Minute())
	//fmt.Println(now.Second())
	////时间戳
	//unix := now.Unix()     //时间戳
	//nano := now.UnixNano() //纳秒的时间戳
	//fmt.Println(unix)
	//fmt.Println(nano)
	//t := time.Unix(1658467085, 0)
	//fmt.Println(t)
	////时间间隔
	//fmt.Println(time.Second)
	////时间相加
	//fmt.Println(now.Add(time.Hour * 24)) //2022-07-23 13:24:59.13339 +0800 CST m=+86400.002217801
	////时间相减
	//now1 := time.Now()
	//fmt.Println(now1.Sub(now).Microseconds())
	//定时器
	//tick := time.Tick(time.Second)
	//for i:=range tick{
	//	fmt.Println(i)
	//}
	//时间格式化 2006 01 02 03 04 05
	//t1 := time.Now()
	//fmt.Println(t1.Format("2006-01-02-03-04-05"))//12小时
	//fmt.Println(t1.Format("2006/01/02/15/04/05")) //24小时
	//fmt.Println(t1.Format("15:04 2006/01/02"))
	//fmt.Println(t1.Format("2006/01/02"))

	//1.获取当前时间，格式化输出为2017/06/19 20:30:05`格式。
	//now := time.Now()
	//fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	////按照对应的格式解析字符串类型的时间
	//parse, err := time.Parse("2006-01-02", "2022-07-22")
	//if err!=nil{
	//	fmt.Printf("err:%v\n",err)
	//	return
	//}
	//fmt.Println(parse)
	//fmt.Println(parse.Unix())
	//now := time.Now()
	//unix := now.Unix()
	//timestampDemo2(unix)
}
func main() {
	//解析字符串时间
	now1 := time.Now()
	fmt.Println(now1)
	time.Parse("2006-01-02-15-05", "2022-07-23-16-46")
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("获取时区时间failed,err:%v", err)
		return
	}
	//按照指定时区解析时间
	time, err := time.ParseInLocation("2006-01-02-15-05", "2022-07-23-16-46", loc)
	if err != nil {
		fmt.Printf("解析时区时间failed,err:%v", err)
		return
	}
	fmt.Println(now1)
	//时间相减
	td := now1.Sub(time)
	fmt.Println(td)
}
