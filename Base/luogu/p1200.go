package main

import "fmt"

//你的飞碟在这儿Your Ride Is Here

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	//使用map初始化大写字母对应的值
	m1 := make(map[rune]int)
	j := 1
	for i := 'A'; i <= 'Z'; i++ {
		m1[i] = j
		j = j + 1
	}
	sum1, sum2 := 1, 1
	//直接遍历字符 然后取map中的值 直接相乘
	for _, v := range s1 {
		sum1 *= m1[v]
	}
	for _, v := range s2 {
		sum2 *= m1[v]
	}
	if sum1%47 == sum2%47 {
		fmt.Println("GO")
	} else {
		fmt.Println("STAY")
	}
}
