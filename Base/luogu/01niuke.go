package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//撤销与恢复
func main() {
	//输入
	input := bufio.NewReader(os.Stdin)
	s, _ := input.ReadString('\n')
	//去除前后端空格
	s1 := strings.TrimSpace(s)
	//分割字符串
	sl := strings.Split(s1, " ")
	stack := []string{}
	redo := []string{}
	for _, v := range sl {
		if v == "redo" {
			if len(redo) > 0 {
				stack = append(stack, redo[len(redo)-1])
				redo = redo[:len(redo)-1]
			}
		} else if v == "undo" {
			if len(stack) > 0 {
				redo = append(redo, stack[len(stack)-1])
				//置为空
				stack = stack[:len(stack)-1]
			}
		} else {
			redo = redo[:0]
			stack = append(stack, v)
		}
	}
	//拼接字符串
	res := strings.Join(stack, " ")
	fmt.Println(res)
}
