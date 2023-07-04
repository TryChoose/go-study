package main

import "fmt"

func main() {
	n := 10
	var f1 float64 = 100.00
	var flag bool = true
	s1 := "王大锤"
	fmt.Printf("n:%d,f1:%f,flag:%v,s1:%s\n", n, f1, flag, s1)

	s2 := "hello沙河小王子"
	count := 0
	r1 := []rune(s2)
	for i := 0; i < len(r1); i++ {
		if !(r1[i] >= 'A' && r1[i] <= 'Z' || r1[i] >= 'a' && r1[i] <= 'z') {
			count++
		}
	}
	fmt.Println("汉字的个数是：", count)

}
