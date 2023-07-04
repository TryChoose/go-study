package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	res := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < len(a); i++ {
		count := 0
		for j := 0; j < i; j++ {
			if a[i] > a[j] {
				count++
			}
		}
		res = append(res, count)
	}
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}
