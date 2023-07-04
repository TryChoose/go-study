package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	l, r := 0, 0
	for _, v := range s {
		if v == ')' {
			r++
		}
		if v == '(' {
			l++
		}
		//判断一下右括号是否先出现 如果是直接NO
		if r > l {
			fmt.Println("NO")
			return
		}
	}
	if l == r {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
