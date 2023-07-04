package main

import (
	"fmt"
)

var (
	up   = 6
	down = 4
	open = 5
	exit = 1
)

func main() {
	var n, x, sum, max int
	var s [100000000]bool
	s[0] = true //0层不开门
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if !s[x] {
			sum += open
		}
		s[x] = true
		if x > max {
			max = x
		}
	}
	sum += n * exit
	sum += max * (up + down)
	fmt.Println(sum)
}
