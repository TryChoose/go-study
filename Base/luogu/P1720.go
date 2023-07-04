package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a [50]int
	a[0], a[1], a[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Printf("%.2f", float64(a[n]))
}
