package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var rep = `[\u4e00-\u9fa5]*`

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
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
		return
	}
}
func getChinese(urL string) {
	pageStr := getURL(urL)
	re := regexp.MustCompile(rep)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}
func main() {
	str := getURL("https://www.ucas.ac.cn/site/74")
	getChinese(str)

}
