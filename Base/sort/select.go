package main

import "fmt"

/*
选择排序
稳定 O(n^2)
*/

func selectSort(a []int) {
	var min int
	for i := 0; i < len(a); i++ {
		min = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
func main() {
	aSli := []int{5, 2, 1, 3, 4, 1, 3, 5, 123, 123, 13, 123}
	selectSort(aSli)
	fmt.Println(aSli)
}
