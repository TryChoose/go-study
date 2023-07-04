package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	x := a*60 + b
	y := c*60 + d
	s := y - x
	fmt.Println(s/60, s%60)
}
