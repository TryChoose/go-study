package main

import "fmt"

//乘积最大
func main() {
	var n, m int //需要m个
	fmt.Scan(&n, &m)
	for i := m; i > 0; i-- {
		fmt.Printf("%d ", n/i)
		n -= n / i
	}
}
