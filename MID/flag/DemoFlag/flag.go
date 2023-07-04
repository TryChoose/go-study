package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "张三", "请输入姓名")
	age := flag.Int("age", 100, "请输入年龄")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}
