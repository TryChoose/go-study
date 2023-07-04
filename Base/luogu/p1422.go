package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0.0
	if n <= 150 {
		sum = float64(n) * 0.4463
	}
	if n >= 151 && n <= 400 {
		sum = float64(n-150)*0.4663 + 150.0*0.4463
	}
	if n >= 401 {
		sum = float64(n-400)*0.5663 + 250*0.4663 + 150*0.4463
	}
	fmt.Printf("%.1f", sum)
}
