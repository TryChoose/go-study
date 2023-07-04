package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}
func main() {
	v := 1
	n := incr(&v)
	fmt.Println(n)
	fmt.Println(incr(&v))
}
