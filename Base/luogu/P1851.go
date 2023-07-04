package main

import "fmt"

func f(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
func main() {
	var s int
	fmt.Scan(&s)
	for i := s; i <= 100000; i++ {
		//先求出s的因子之和
		x := f(i)
		//求出x的因子之和
		y := f(x)
		//画个图 对应一把 就可以清楚
		if i == y && i != x {
			fmt.Printf("%d %d", i, x)
			return
		}
	}
}
