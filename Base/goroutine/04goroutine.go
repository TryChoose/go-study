package main

import (
	"fmt"
	"time"
)

// 通过channel

func main() {
	strChan := make(chan int, 1)
	numChan := make(chan int, 1)
	strChan <- 0 // 很关键 str管道输入一个值之后 这个管道是处于一个阻塞状态,只有释放之后才能进行下一步从而实现并行
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			<-strChan
			fmt.Print(string(i) + " ")
			numChan <- int(i)
		}
	}()

	go func() {
		for i := 0; i < 26; i++ {
			<-numChan
			fmt.Print(i+1, " ")
			strChan <- i
		}
	}()
	time.Sleep(time.Millisecond * 50)
	fmt.Println()
}
