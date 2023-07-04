package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./filestat.go")
	if err != nil {
		fmt.Println("open file failed，err:%v", err)
		return
	}
	fmt.Printf("%T\n", fileObj)
	//查询文件的信息
	fileInfo, err := fileObj.Stat()
	fmt.Println(fileInfo.Name(), fileInfo.Size())

}
