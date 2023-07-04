package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := []rune{}
	for _, v := range s {
		if v >= '0' && v <= '9' {
			a = append(a, v)
		}
	}
	for _, v := range a {
		fmt.Printf("%c ", v)
	}

}
