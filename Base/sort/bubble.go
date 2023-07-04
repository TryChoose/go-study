package main

import "fmt"

/*
冒泡排序
稳定算法 O(n^2)
*/
func bubble(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
func main() {
	aSli := []int{5, 2, 1, 3, 4}
	bubble(aSli)
	fmt.Println(aSli)
}
