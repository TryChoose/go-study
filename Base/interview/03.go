package main

import "fmt"

//数组去重
func remove(list []int) []int {
	var x = []int{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
func reCount(arr []int) int {
	res := make([]int, len(arr))
	count := 0
	for _, v := range arr {
		res[v]++
	}
	fmt.Println(res)
	for _, v := range res {
		if v >= 2 {
			count++
		}
	}
	return count
}
func main() {
	a := []int{1, 2, 3, 4, 4, 3, 2, 1, 2, 4, 5, 1, 0, 9, 2, 3}
	//i := remove(a)
	//fmt.Println(i)
	//fmt.Println(a)
	b := reCount(a)
	fmt.Println(b)
}
