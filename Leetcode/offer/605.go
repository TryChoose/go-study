package main

import "fmt"

//
// main
//  @Description: 程序开启
//
func main() {
	flowers := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
	fmt.Println(flowers)
}

//
// canPlaceFlowers
//  @Description: 种花问题
//  @param flowerbed
//  @param n
//  @return bool
//
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	t := 1
	for _, v := range flowerbed {
		if v == 0 {
			t++
		} else {
			count += (t - 1) / 2
			t = 0
		}
	}
	count += t / 2
	fmt.Println(count)
	if count >= n {
		return true
	}
	return false
}
