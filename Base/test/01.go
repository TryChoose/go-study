package main

import "fmt"

// 最大连续1的个数
//描述
//给定一个二进制数组，找出该数组中最大连续1的个数。
func Cmp(a int, b int) int {
	if a > b {
		b = a
	}
	return b
}
func FindMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			max = Cmp(max, count)
			count = 0
		}
	}
	max = Cmp(max, count) //防止最后面都是1 无法进入else 语句 无法进行比较
	return max
}
func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(FindMaxConsecutiveOnes(nums))
}
