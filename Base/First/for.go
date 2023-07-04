package main

import "fmt"

/*
 for循环
 第一类型：for init; condition; post { }
*/
func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
