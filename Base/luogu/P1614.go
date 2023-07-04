package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n == 0 {
		fmt.Println(0)
		return
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	k := 0
	min := 999999999
	for i := 0; i < n; i++ {
		sum := 0
		k = 0
		for j := i; j < i+m && j < n; j++ {
			sum += a[j]
			k++
		}
		if min > sum && k == m {
			min = sum
		}
	}
	fmt.Println(min)
}
