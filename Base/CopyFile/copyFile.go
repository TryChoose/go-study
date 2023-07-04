package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//将一个文件里面的内容写入到另外一个文件里面

func copyFile(dstName string, srcName string) {
	//1.需要读取源文件的内容
	//1.1打开源文件
	open, err := os.Open(srcName)
	//1.2记得关闭源文件
	defer open.Close()
	//1.3判断是否有错误
	if err != nil {
		fmt.Printf("打开文件有错误：%v", err)
	}
	//使用了ioutil.ReadFile(srcName)获取文件
	file, err := ioutil.ReadFile(srcName)
	if err != nil {
		fmt.Printf("读取文件有错误:%v", err)
	}
	//2.将源文件的内容写入到dst里面
	//2.1直接采用ioutil.WriteFile读取
	err1 := ioutil.WriteFile(dstName, []byte(file), 0666)
	//2.2判断是否有错误
	if err1 != nil {
		fmt.Printf("写入文件有错误:%v", err1)
	}
	fmt.Println("copy 成功")
}
func main() {
	copyFile("02", "01")
}
