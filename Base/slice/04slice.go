package main

import "fmt"

//元素旋转

/*
示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
*/
func rotate(slice []int, k int) []int {
	//1.取出最后一个元素
	for i := 0; i < k; i++ {
		x := slice[len(slice)-1]
		//fmt.Println(x)
		//2.移动除去最后一个元素的往后移一位
		copy(slice[1:], slice[0:len(slice)-1])
		//3.将最后一个元素放在第一位
		slice[0] = x
	}

	return slice
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotate(slice, 3))
}
