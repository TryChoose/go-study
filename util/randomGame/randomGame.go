package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机生成一个1-100（不包含）作为目标值
随机生成20个作为候选值
*/
func candidateNums() (res []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		n := rand.Intn(100)
		res = append(res, n)
	}
	return
}
func targetNum() int {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)
	return target
}
func main() {
	target := targetNum()
	fmt.Println("中奖号码：", target)
	time.Sleep(1)
	count := 0
	nums := candidateNums()
	fmt.Println("候选号码：")
	for _, v := range nums {
		fmt.Printf("%d ", v)
		if target == v {
			count++
		}
	}
	fmt.Println()
	fmt.Println("中了多少个：", count)
}
