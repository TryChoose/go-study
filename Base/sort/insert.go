package main

import "fmt"

/*
插入排序
稳定 O(n^2)
*/
func insert(a []int) {
	var j int
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t := a[i]
			a[i] = a[i-1]
			for j = i - 1; j >= 0 && a[j] > t; j-- {
				a[j+1] = a[j]
			}
			a[j+1] = t
		}
	}
}
func main() {
	aSli := []int{5, 2, 1, 3, 4}
	insert(aSli)
	fmt.Println(aSli)
}
