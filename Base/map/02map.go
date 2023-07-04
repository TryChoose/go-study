package main

import "fmt"

func main() {
	m1 := map[int]int{
		0: 2,
		1: 7,
		2: 11,
		3: 15,
	}
	fmt.Println(m1[0])      //值
	if _, ok := m1[1]; ok { //ok是bool值,p应该是索引
		fmt.Println(ok)
	}
}
