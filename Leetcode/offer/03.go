package main

import "fmt"

func f1(n, x int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2 * n
	}
	return 2*x*f1(n-1, x) - 2*(n-1)*f1(n-2, x)
}
func main() {
	var n, x int
	fmt.Scan(&n, &x)
	fmt.Println(f1(n, x))
}
