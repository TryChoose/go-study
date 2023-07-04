package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

/*
1.先寻找网站
2.将网站的body转为字符串并且返回
3.通过正则表达式提取我们需要的内容
4.然后在通过downFile下载
*/
var (
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(gif)))`
)

//爬取网站的图片地址
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
	//返回一个string类型 整个页面
}

//获取所有图片的网址
func getImag(urL string) {
	pageStr := getURL(urL)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	index := 0
	for _, result := range results {
		fmt.Printf("%T,%v\n", result[0], result[0])
		//去头去尾"
		//trimPrefix := strings.TrimPrefix(result[0], "\"")
		//suffix := strings.TrimSuffix(trimPrefix, "\"")
		//fmt.Println(suffix)
		i := strconv.Itoa(index)
		downFile(result[0], i+".jpg")
		index++
		//fmt.Println(index)
	}
}

//处理错误
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
		return
	}
}

//下载全部的图片
func downFile(url string, fileName string) bool {
	pageStr := getURL(url)
	fileName = "E:\\AStudyGoland\\src\\csgo\\MID\\reptile\\02reptile\\image\\" + fileName
	//写出数据
	err := ioutil.WriteFile(fileName, []byte(pageStr), 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}
func main() {
	//getImag("https://www.zcool.com.cn/")

	getImag("https://max.book118.com/html/2022/0703/8027052011004115.shtm")
	//getImag("https://www.renrendoc.com/paper/224801415.html")
	//downFile("https://view-cache.book118.com/view25/M03/36/32/wKh2HWLFP8SAa_YhAAdJeg2KloA474.png", "3.jpg")
}
