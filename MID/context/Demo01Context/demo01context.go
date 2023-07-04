package main

import (
	"fmt"
	"sync"
	"time"
)

/*
使用传统方式关闭子 goroutine
*/
var wg sync.WaitGroup
var lock sync.Mutex
var i int

func f1() {
	defer wg.Done()
	for {
		fmt.Println(i)
		lock.Lock()
		i++
		lock.Unlock()
		time.Sleep(time.Millisecond)
		if i == 10 {
			break
		}
	}
}
func main() {
	wg.Add(1)
	go f1()
	time.Sleep(time.Millisecond * 10)
	wg.Wait()
}
