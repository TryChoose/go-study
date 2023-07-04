package main

import "fmt"

func f1() {
	//求平均年龄
	var n, age int
	fmt.Scan(&n)
	t := n
	avg := 0.0
	for n > 0 {
		fmt.Scan(&age)
		avg += float64(float64(age) / float64(t))
		n--
	}
	fmt.Printf("%.2f", avg)
}
func main() {
	f1()
}
