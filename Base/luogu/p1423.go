package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	s, sum, count := 2.0, 2.0, 1
	for {
		s = s * 0.98
		count++
		sum += s
		if sum > n {
			break
		}
	}
	fmt.Println(count)
}
