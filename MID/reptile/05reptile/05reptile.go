package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var (
	//提取答案链接 正则表达式
	ansReg    = `<a class="question-link" href="/question/(\d)+`
	title     = `<div class="question-stem">.+</div>`
	queClass  = `<div class="question-tx">.+</div>`
	queOptReg = `<div class="option-item">.+</div>`
	queAnsOpt = `<div class="question-answer">.+</div>`
	ss        = []string{"", "学历类", "职业资格类", "外语类", "计算机类", "财会类", "建筑类", "医药类", "外贸类", "公务员类", "考研类", "趣味测试类"}
	s1        string
	fileIndex int
)

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}
func getWebURL(s1 string) (str string) {
	//获取网站
	resp, err := http.Get(s1)
	HandleError(err, "获取URL失败！")
	//读取页面
	pageByte, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "获取页面信息失败！")
	//字节转为字符
	pageStr := string(pageByte)
	//fmt.Println(pageStr)
	return pageStr
}

// 提取分类个数
func categorize(start, end int) {
	for i := start; i <= end; i++ {
		if i == 6 || i == 8 || i == 11 {
			continue
		}
		index := 1
		//byBuFio(ss[i])
		for {
			str := getWebURL("https://www.woaishuati.com/album/" + strconv.Itoa(i) + "?" + "page=" + strconv.Itoa(index))
			reg := regexp.MustCompile(ansReg)
			res := reg.FindAllStringSubmatch(str, -1)
			for _, v := range res {
				b := []byte(v[0])
				b = b[31:]
				s1 = "https://www.woaishuati.com" + string(b)
				class := questionClass(s1)
				//fmt.Println(class)
				//questionTitle(s1)
				//if class == "【填空题】" || class == "【简答题】" {
				//	questionAns(s1)
				//	continue
				//}
				//questionOpt(s1)
				//questionAns(s1)
				if class == "简答题" {
					break
				}
				questionTitle(s1)
				if class == "填空题" {
					questionAns(s1)
					continue
				}
				questionOpt(s1)
				questionAns(s1)
			}
			if index == 30 {
				break
			}
			index++
		}
	}
}

// 提取题目类型
func questionClass(str string) string {
	queCls := getWebURL(str)
	reg := regexp.MustCompile(queClass)
	res := reg.FindAllStringSubmatch(queCls, 1)
	b := res[0][0]
	b = b[28 : len(b)-9]
	byBuFio(string(b))
	return string(b)
}

// 提取题干
func questionTitle(str string) {
	//fmt.Println(str)
	titleURL := getWebURL(str)
	reg := regexp.MustCompile(title)
	res := reg.FindAllStringSubmatch(titleURL, 1)
	//b := res[0][0]
	b := res[0][0]
	b = b[27 : len(b)-6]
	byBuFio(string(b))
}

// 提取选项
func questionOpt(str string) {
	queOpt := getWebURL(str)
	reg := regexp.MustCompile(queOptReg)
	res := reg.FindAllStringSubmatch(queOpt, -1)
	countA := 0
	//fmt.Println(res)
	for _, v := range res {
		b := []byte(v[0])
		b = b[25 : len(b)-6] //选项
		if b[0] == 'A' {
			countA++
		}
		if countA == 0 {
			questionAns(s1)
			break
		}
		//fmt.Println(b)
		if countA == 2 {
			countA = 0
			break
		}
		byBuFio(string(b))
	}
}

// 提取答案
func questionAns(str string) {
	queAnsOption := getWebURL(str)
	reg := regexp.MustCompile(queAnsOpt)
	res := reg.FindAllStringSubmatch(queAnsOption, 1)
	//for _, v := range res {
	//	b := []byte(v[0])
	//	b = b[29 : len(b)-6]
	//	//byBuFio("答案:" + string(b))
	//	byBuFio(string(b))
	//}
	b := res[0][0]
	b = b[29 : len(b)-6]
	//fmt.Println(string(b)) //答案
	byBuFio("答案:" + string(b))
	//fmt.Println(string(b)) //答案

}
func byBuFio(s string) {
	file, err := os.OpenFile("0"+strconv.Itoa(fileIndex)+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("write file faiied ：%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(s + "\n")
	writer.Flush()
}

func main() {
	//fmt.Println("分类起始值必须为1")
	//categorize(1, 11)
	for fileIndex = 1; fileIndex <= 11; fileIndex++ {
		if fileIndex == 3 || fileIndex == 6 || fileIndex == 8 || fileIndex == 11 {
			continue
		}
		categorize(fileIndex, fileIndex)
	}
}
