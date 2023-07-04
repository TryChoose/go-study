package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件操作
//自定义读取的字符的个数
func readFileLowMethod1() {
	//打开要读去的文件
	openFile, err := os.Open("./01File.go")
	if err != nil {
		fmt.Printf("open file faild ：%v", err)
	}
	//关闭文件
	defer openFile.Close()
	//读取该文件
	var fileByte [128]byte
	for {
		n, err := openFile.Read(fileByte[:])
		if err == io.EOF {
			return

		}
		if err != nil {
			fmt.Printf("read file faild %v", err)
			return
		}

		fmt.Println("读取了%d个字节:", n)
		fmt.Println(string(fileByte[:n]))
		if n < 128 {
			return
		}
	}
}

//单行读取
func bufIoRead() {
	//打开要读的文件
	openFile, err := os.Open("./01File.go")
	if err != nil {
		fmt.Printf("读取出现了错误：%v", err)
		return
	}
	defer openFile.Close()
	//创建一个可以读的对象
	reader := bufio.NewReader(openFile)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed %v", err)
			return
		}
		fmt.Print(line)
	}
}

//读取整个文件
func iOutFile() {
	ret, err := ioutil.ReadFile("./01File.go")
	if err != nil {
		fmt.Printf("read file failed%v", err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	//readFileLowMethod1()
	//bufIoRead()
	iOutFile()
}
