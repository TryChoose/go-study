package main

import (
	"fmt"
	"unicode"
)

//判断字母，数字，其他符号的数量
func charCount(s string) (x, y, z int) {
	r1 := []rune(s)
	for _, v := range r1 {
		if unicode.IsLetter(v) {
			x++
		} else if unicode.IsNumber(v) {
			y++
		} else {
			z++
		}
	}
	return
}
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(charCount(s))
}
