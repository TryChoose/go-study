package main

import (
	"fmt"
	"unicode"
)

//判断字符串里面的汉字个数
func main() {
	var s1 string
	count := 0
	fmt.Scanln(&s1)
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}
