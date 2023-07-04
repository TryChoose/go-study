package main

import "fmt"

// 和与平均值
func main() {
	n := 0
	sum := 0.0
	avg := 0.0
	fmt.Scan(&n)
	a := make([]int, n)
	b := []int{}
	for i := 0; i < int(n); i++ {
		fmt.Scan(&a[i])
		sum += float64(a[i])
		b = append(b, a[i])
	}
	avg = sum / float64(n)
	fmt.Printf("%.0f %.5f", sum, avg)
}
