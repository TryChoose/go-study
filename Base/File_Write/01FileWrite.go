package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//Write和Writing方法

func write01() {
	file, err := os.OpenFile("01.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("write file faiied ：%v", err)
		return
	}
	defer file.Close()
	str := "张三李四王五二麻子\n"
	file.Write([]byte(str))
	file.WriteString("指数上的上的")
}
func byBuFio() {
	file, err := os.OpenFile("01.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		fmt.Printf("write file faiied ：%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("漳卅你无\n")
	writer.Flush()
}
func iouTiFile() {
	str := "张三为的身上的"
	err := ioutil.WriteFile("01.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file faiied ：%v", err)
		return
	}
}
func keyWrite() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}
func main() {
	//Write01()
	byBuFio()
	//iouTiFile()
	//keyWrite()
}
