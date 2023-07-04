package main

import (
	"fmt"
	"time"
)

// 不管数组怎么变化 始终保持f1 f2 f3
var ch1 = make(chan int, 1)
var ch2 = make(chan int, 1)
var ch3 = make(chan int, 1)

func f1() {
	<-ch1
	fmt.Println("First")
	ch2 <- 0
}
func f2() {
	<-ch2
	fmt.Println("Second")
	ch3 <- 0
}
func f3() {
	<-ch3
	fmt.Println("Third")
}
func main() {
	ch1 <- 0
	time.Sleep(1 * time.Second)
	go f1()
	time.Sleep(1 * time.Second)
	go f2()
	time.Sleep(1 * time.Second)
	go f3()
	time.Sleep(1 * time.Second)
}
