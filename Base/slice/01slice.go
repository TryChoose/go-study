package main

import "fmt"

//reverse 函数

//slice版本

//切片传递是地址 ,是内部的底层数组
func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	//return s
}

//数组指针
func reversePoint(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	//s1:=[5]int{2,4,6,8,10}
	//reverseSlice(s)
	//fmt.Println(s)
	//reversePoint(&s1)
	//fmt.Println(s1)
	//将元素做移动两位
	reverseSlice(s[:3])
	reverseSlice(s[3:])
	reverseSlice(s)
	fmt.Println(s)
}
