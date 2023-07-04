package main

import (
	"fmt"
)

//反转

func main() {
	a := []int{}
	var n int
	var i int = 0
	fmt.Scan(&n)
	for n != 0 {
		a = append(a, n)
		i++
		fmt.Scan(&n)
	}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}
