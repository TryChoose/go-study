package main

import "fmt"

//开灯
func main() {
	var a [2000000]int
	//次数
	var n, op int
	fmt.Scan(&n)
	var f float64
	//输入实数 和整数
	for i := 0; i < n; i++ {
		fmt.Scan(&f, &op)
		for j := 1; j <= op; j++ {
			//判断开关灯
			if a[int(float64(j)*f)] == 0 {
				a[int(float64(j)*f)] = 1
			} else {
				a[int(float64(j)*f)] = 0
			}
		}
	}
	for i := 0; ; i++ {
		if a[i] == 1 {
			fmt.Println(i)
			break
		}
	}
}
