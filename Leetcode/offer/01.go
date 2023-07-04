package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

//版本号
//分割字符串母字符串
//分割子字符串
//排序

func main() {
	input := bufio.NewReader(os.Stdin)
	readString, _ := input.ReadString('\n')
	//fmt.Println(readString)
	str := strings.TrimSpace(readString)
	split := strings.Split(str, "#")
	//fmt.Println(split)
	sort.Strings(split)
	for _, v := range split {
		fmt.Println(v)
	}
}
