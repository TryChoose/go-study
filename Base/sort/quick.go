package main

import "fmt"

/*
快速排序
不稳定,O(nLogN)
*/

func quick(a []int, l, r int) {
	if l < r {
		m := part(a, l, r)
		quick(a, l, m-1)
		quick(a, m+1, r)
	}
}

// 进行分区  返回中枢值索引
func part(a []int, l, r int) int {
	k := a[r]
	i, j := l, l
	for j < r {
		if a[j] < k {
			a[i], a[j] = a[j], a[i]
			i++
		}
		j++
	}
	a[i], a[r] = a[r], a[i]
	return i
}

func main() {
	aSli := []int{54, 73, 21, 35, 67, 78, 63, 24, 89}
	quick(aSli, 0, len(aSli)-1)
	fmt.Println(aSli)
}
