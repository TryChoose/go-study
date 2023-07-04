package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
在一个文件内容里面单独插入新的内容
*/

func InsertFile(src string, des string) {
	/*
		1.分析先需要打开src根据需求把光标移动到插入的位置的中间
			1.1 现需要打开源文件
		2.读取光标之前的内容，在一个临时文件里面,
		3.在这个临时文件里面插入题目的所要求的内容
		4.将剩下的源文件的内容读取到临时文件当中
		5.最后清空源文件,把临时文件的内容copy到源文件中
	*/
	open, err := os.Open(src)
	if err != nil {
		fmt.Printf("打开源文件出错了！！ 出错的内容是：%v", err)
		return
	}
	defer open.Close()
	//读取文件
	reader := bufio.NewReader(open)
	//写入文件
	WriteFile, err1 := os.OpenFile(des, os.O_CREATE|os.O_APPEND, 066)
	if err1 != nil {
		fmt.Printf("写入文件出错！：%v", err1)
	}
	defer WriteFile.Close()
	for i := 1; i <= 20; i++ {
		line1, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取文件出错了！！ 读取的文件出错了:%v", err)
			return
		}
		fmt.Print(line1)
		writer := bufio.NewWriter(WriteFile)
		writer.WriteString(line1)
		writer.Flush()
	}
	var s = "-------------------------------\n"
	writer := bufio.NewWriter(WriteFile)
	writer.WriteString(s)
	writer.Flush()
	for {
		line1, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("读取文件出错了！！ 读取的文件出错了:%v", err)
			return
		}
		fmt.Print(line1)
		writer := bufio.NewWriter(WriteFile)
		writer.WriteString(line1)
		writer.Flush()
	}
}
func f1(src string, des string, k int) {
	//在一个句子里面插入一句话
	//先将要插入的位置前的字符读取在另外一个文件当中,然后在将这个文件插入需要的内容，在读取源文件里面的剩下的内容
	//打开原始文件
	fileSrc, err := os.OpenFile(src, os.O_RDWR, 055)
	if err != nil {
		fmt.Printf("打开文件失败：err：", err)
		return
	}
	defer fileSrc.Close()
	//打开目标文件
	fileDes, err1 := os.OpenFile(des, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err1 != nil {
		fmt.Printf("faild：err1：%v", err1)
		return
	}
	defer fileDes.Close()
	//循环读写
	for i := 0; i < k; i++ {
		//开始读取原始文件的内容
		var context = [1]byte{}
		n, err2 := fileSrc.Read(context[:])
		if err2 != nil {
			fmt.Printf("读取文件失败：err2:%v", err2)
			return
		}
		//写入文件
		fileDes.Write(context[:n])
	}
	//写入新的内容
	var s []byte
	s = []byte{'2', '3', '4'}
	fileDes.Write(s)
	//把剩下原来的字符写入新的文件
	//os实现了seek方法
	var x [1024]byte
	for {
		n, err := fileSrc.Read(x[:])
		if err == io.EOF {
			fileDes.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("读取错误！err:%v", err)
		}
		fileDes.Write(x[:n])
		fmt.Println(n)
	}
	//os.Rename(des,"021")
}
func main() {
	//InsertFile("01", "02")
	f1("011", "022", 2)
}
