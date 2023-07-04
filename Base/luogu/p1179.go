package main

import "fmt"

//2出现的次数
func main() {
	var l, r int
	count := 0
	fmt.Scan(&l, &r)
	for i := l; i <= r; i++ {
		n := i
		for n != 0 {
			if n%10 == 2 {
				count++
			}
			n = n / 10
		}
	}
	fmt.Println(count)
}
