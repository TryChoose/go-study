package main

import "fmt"

// 数组形式的整数加法
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	//return s
}
func addToArrayForm(n []int, m int) []int {
	a := []int{}
	for i := len(n) - 1; i >= 0; i-- {
		sum := n[i] + m%10
		m /= 10
		if sum >= 10 {
			m++
			sum -= 10
		}
		a = append(a, sum)
	}
	for ; m > 0; m /= 10 {
		a = append(a, m%10)
	}
	reverse(a)
	return a
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(addToArrayForm(a, 24))

}
