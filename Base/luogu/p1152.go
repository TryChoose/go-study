package main

import (
	"fmt"
	"sort"
)

//欢乐的跳
func main() {
	var n int
	var flag bool = true
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 1; i < n; i++ {
		b[i-1] = a[i] - a[i-1]
	}
	sort.Ints(b)
	for i := 1; i <= n-1; i++ {
		if b[i] > n-1 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("Jolly")
	} else {
		fmt.Println("Not jolly")
	}
}
