package main

import "fmt"

//开灯问题
func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i*i <= n; i++ {
		fmt.Printf("%d ", i*i)
	}
}
