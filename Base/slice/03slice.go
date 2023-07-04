package main

import "fmt"

//利用copy函数删除该元素并且移动
//移除元素 并且顺序 相当于 重新赋值
func remove1(slice []int, i int) []int {
	//fmt.Println(slice[i:],slice[i+1:]) //[5 8 9] [8 9]
	copy(slice[i:], slice[i+1:])
	//fmt.Println(slice[i:],slice[i+1:]) //[8 9 9] [9 9]
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
func main() {
	s := []int{2, 4, 5, 8, 9}
	fmt.Println(remove1(s, 2))
	s2 := []int{2, 4, 5, 8, 9}
	fmt.Println(remove2(s2, 2))
}
