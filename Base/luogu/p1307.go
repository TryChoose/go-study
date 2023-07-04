package main

import "fmt"

func main() {
	var n int
	s := 0
	fmt.Scan(&n)
	for n != 0 {
		s = s*10 + n%10
		n = n / 10
	}
	fmt.Println(s)

}
