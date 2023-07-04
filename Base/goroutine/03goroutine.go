package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int = 0

// 不要通过共享内存来通信，而应该通过通信来共享内存
// 通过锁来实现
func Count01(lock *sync.Mutex) {
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
}
func main() {
	//runtime.Gosched()作用是让当前goroutine让出CPU，
	//好让其它的goroutine获得执行的机会。同时，当前的goroutine也会在未来的某个时间点继续运行。
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count01(lock)
	}
	for {
		lock.Lock()
		c := count
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
