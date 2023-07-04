package main

import (
	"fmt"
	"sync"
	"time"
)

/*
使用通道方式关闭子goroutine
注意需要初始化管道 要不然会一直堵塞
*/
var wg sync.WaitGroup
var lock sync.Mutex
var i int
var exitChan chan bool = make(chan bool, 1)

func f1() {
	defer wg.Done()
FLoop:
	for {
		fmt.Println(i)
		lock.Lock()
		i++
		lock.Unlock()
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break FLoop
		default:
		}
	}
}
func main() {
	wg.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	exitChan <- true
	close(exitChan)
	wg.Wait()
}
