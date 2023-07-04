package main

import "fmt"

//压缩技术
func main() {
	var n, b int
	fmt.Scan(&n)
	a := make([]int, n*n)
	var i = 0
	count := 1 //判断奇偶性
	for {
		//奇数循环输出0,偶数输出1 ,可以存入a
		fmt.Scan(&b)
		if count%2 == 1 {
			for j := 0; j < b; j++ {
				a[i] = 0
				i++
			}
		} else if count%2 == 0 {
			for j := 0; j < b; j++ {
				a[i] = 1
				i++
			}
		}
		if n*n == i { //循环退出条件
			break
		}
		count++
	}
	j := 0
	for i := 0; i < n*n; i++ {
		fmt.Printf("%d", a[i])
		j++
		if j == n {
			j = 0
			fmt.Println()
		}
	}
}
