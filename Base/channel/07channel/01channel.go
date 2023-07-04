/**
 * @Time: 2023/6/20 9:33
 * @Author: yncn90@gmail.com
 * @File: 01channel.go 01channel
 * @Software: GoLand main.go
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// printNum 单个处理器顺序打印数字
func printNum() {
	runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	wg.Add(20)
	ch1 := make(chan int, 1)
	for i := 0; i < 10; i++ {
		ch1 <- i
		go func() {
			fmt.Println("i=", <-ch1)
			defer wg.Done()
		}()

	}
	for i := 0; i < 10; i++ {
		ch1 <- i
		go func(ch1 <-chan int) {
			select {
			case v := <-ch1:
				fmt.Println("v=", v)
			default:
				fmt.Println("default")
			}
			defer wg.Done()
		}(ch1)
	}
	defer close(ch1)
	wg.Wait()
}

var (
	count int
	rw    = new(sync.RWMutex)
	wg    sync.WaitGroup
)

func read() {
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			rw.RLock() //读锁
			fmt.Println(count)
			rw.RUnlock()

		}
	}()
}
func write() {
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			rw.Lock() //写锁
			count++
			rw.Unlock()
		}
	}()
}
func main() {
	wg.Add(2)
	write()
	read()
	wg.Wait()

	printNum()
}
