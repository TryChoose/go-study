package main

import "fmt"

// 多携程计算和
var ch = make(chan int, 2)

func sum(ch chan int, a []int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	ch <- sum
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go sum(ch, a[:len(a)/2])
	go sum(ch, a[len(a)/2:])
	s1, s2 := <-ch, <-ch
	fmt.Println(s1 + s2)
	fmt.Println(s1 + s2)

}
