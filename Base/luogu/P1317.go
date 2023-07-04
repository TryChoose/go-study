package main

import "fmt"

//低洼
func main() {
	var n int
	var count = 0
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < n-1; i++ {
		if a[i-1] >= a[i] && a[i] < a[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
