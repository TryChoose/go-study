package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁

var (
	x      int
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func readWithRWLock() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}
func writeWithRWLock() {
	rwLock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go writeWithRWLock()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go readWithRWLock()
		wg.Add(1)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x, end.Sub(start)) //10 172.4488ms
}
