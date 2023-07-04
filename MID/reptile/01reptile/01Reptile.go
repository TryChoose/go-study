package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
爬虫：
1.爬取邮箱
*/

var (
	regQQ    = `[a-z]+(\d+|[a-z]+)*@\d+.com`
	regEmail = `\w+@\w+(\.\w+)?`
	//+正闭包
	//\s\S各种字符
	//+?代表贪婪模式
	relink   = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIDCard = `([1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx])`
	//reImg    = `"https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))"`

	reImg2 = `"https?://[^"]+?(\.((jpg)|(png)))"`
)

func getURL(s1 string) (pageStr string) {
	//获取网站
	resp, err := http.Get(s1)
	HandleError(err, "获取URL失败！")
	//读取页面
	pageByte, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "获取页面信息失败！")
	//字节转为字符
	pageStr = string(pageByte)
	return pageStr
}
func getEmail(url string) {
	pageStr := getURL(url)
	//fmt.Println(pageStr)
	//过滤数据，过滤qq邮箱
	re := regexp.MustCompile(regEmail)
	//-1代表全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)
	for _, result := range results {
		fmt.Println("邮箱：", result[0])
	}
}
func getLink(url string) {
	pageStr := getURL(url)
	re := regexp.MustCompile(relink)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("地址：", result[0])
	}
}
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
		return
	}
}
func getTel(urL string) {
	pageStr := getURL(urL)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("电话号：", result[0])
	}
}
func getIDCard(urL string) {
	pageStr := getURL(urL)
	re := regexp.MustCompile(reIDCard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("身份号：", result[0])
	}
}
func getImag(urL string) {
	pageStr := getURL(urL)
	re := regexp.MustCompile(reImg2)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}
func main() {
	//getEmail("https://zhidao.baidu.com/question/177832050.html")
	//getLink("http://zhidao.baidu.com/question/1732665224270614947.html")
	//getTel("http://www.tiantianxieye.com/more.php")
	//getIDCard("http://liaocheng.dzwww.com/lcxw/202009/t20200902_6516416.htm")
	getImag("https://www.zcool.com.cn/")
}
