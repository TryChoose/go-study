package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	count := 0
	max := 0
	s := 0
	for i := 1; i < n; i++ {
		if a[i] == a[i-1]+1 {
			count++
			s = count
		} else if a[i] != a[i-1]+1 {
			count = 0
		}
		if max < s {
			max = s
		}
	}
	fmt.Println(max + 1)
}
