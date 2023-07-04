package main

import "fmt"

//回文数

func main() {
	s := "中12321中"
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}
	for i := 0; i < len(s)/2; i++ {
		if r[i] != r[len(r)-i-1] {
			fmt.Println("不是回文")
			break
		}
	}
	fmt.Println("是回文")

}
