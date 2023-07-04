package main

import "fmt"

func main() {
	var x, n int
	sum := 1
	fmt.Scan(&x, &n)
	for i := 0; i < n; i++ {
		sum = sum + sum*x //一个怪兽一轮传染x个怪兽,第二轮就该有x+1个怪兽继续传染
	}
	fmt.Println(sum)
}
