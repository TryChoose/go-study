package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][2]int, n)
	for i := 0; i < len(a); i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	res, sum, b := 0, 0, 0
	for i := 0; i < len(a); i++ {
		sum = 0
		for j := 0; j < 2; j++ {
			sum += a[i][j]
		}
		b += sum - 8
		res += b
	}
	fmt.Println(res)
}
