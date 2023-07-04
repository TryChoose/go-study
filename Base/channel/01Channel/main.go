package main

import (
	"fmt"
	"sync"
)

var ch1 chan int
var wg sync.WaitGroup

func noBufChannel() {
	wg.Add(1)
	ch1 = make(chan int)
	go func() {
		x := <-ch1
		fmt.Println(x)
		wg.Done()
	}()
	ch1 <- 10
	wg.Wait()
}
func bufChannel() {
	//开辟空间
	ch1 = make(chan int, 10)
	//发送
	ch1 <- 10
	//接收
	x := <-ch1
	fmt.Println(x)
	//关闭
	close(ch1)
}
func main() {
	bufChannel()
	noBufChannel()
}
