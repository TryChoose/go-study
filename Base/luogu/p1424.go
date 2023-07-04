package main

import "fmt"

func main() {
	var x, n int
	sum := 0
	fmt.Scan(&x, &n)
	for i := 1; i <= n; i++ {
		if x != 6 && x != 7 {
			sum += 250
		}
		if x == 7 {
			x = 1
		} else {
			x++
		}
	}
	fmt.Println(sum)
}
