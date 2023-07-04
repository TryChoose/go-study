package main

import (
	"fmt"
)

// 子数整数
func main() {
	var k int
	var flag bool
	fmt.Scan(&k)
	for i := 10000; i <= 30000; i++ {
		var s1, s2, s3 int
		var a, b, c, d, e int
		a = i % 10         //个
		b = i / 10 % 10    //十
		c = i / 100 % 10   //百
		d = i / 1000 % 10  //千
		e = i / 10000 % 10 //万
		s1 = e*100 + d*10 + c
		s2 = d*100 + c*10 + b
		s3 = c*100 + b*10 + a
		if s1%k == 0 && s2%k == 0 && s3%k == 0 {
			flag = true
			fmt.Println(i)
		}
	}
	if flag == false {
		fmt.Println("No")
	}
}
