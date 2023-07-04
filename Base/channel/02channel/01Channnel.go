package main

import (
	"fmt"
	"sync"
)

var ch1 chan int

var ch2 chan int

var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2() {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	////并发操作 耗时间s 隔开了一位
	//for x:=range ch1{
	//	_, ok := <-ch1
	//	if !ok {
	//		break
	//	}
	//	ch2 <- x*x
	//
	//}
	close(ch2)
}
func main() {
	wg.Add(2)
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	go f1()
	go f2()
	for i := range ch2 {
		fmt.Println(i)
	}
	wg.Wait()
}
