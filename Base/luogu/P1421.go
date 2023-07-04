package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var c int
	c = a*10 + b
	fmt.Printf("%d", c/19)
}
