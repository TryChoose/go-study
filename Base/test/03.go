package main

import (
	"fmt"
	"sort"
)

// 有序数组的平方
func SquareArray(a []int) []int {
	// write your code here

	//fmt.Println(a)
	for i := 0; i < len(a); i++ {
		a[i] *= a[i]
	}
	sort.Ints(a)
	return a
}
func main() {
	a := []int{-4, -1, 0, 3, 10}
	SquareArray(a)
	fmt.Println(a)
}
