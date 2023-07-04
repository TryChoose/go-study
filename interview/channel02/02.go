package main

import (
	"fmt"
	"sync"
)

// 实现自旋锁
// var Count int32 = 1
//
//	func SpinLock(i int32, fn func()) {
//		for {
//			if n := atomic.LoadInt32(&Count); n == i {
//				fn()
//				atomic.AddInt32(&Count, 1)
//				break
//			}
//			time.Sleep(time.Nanosecond)
//		}
//	}
//
//	func main() {
//		for i := int32(1); i < 4; i++ {
//			go func(i int32) {
//				fn := func() {
//					fmt.Println(i)
//				}
//				SpinLock(i, fn)
//			}(i)
//		}
//		SpinLock(4, func() {})
//	}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch = make(chan int, 3)
	go func() {
		defer wg.Done()
		for i := 1; i < 4; i++ {
			ch <- i
		}
		close(ch) //很关键的 必须要加 要不然就会判断为死锁
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	wg.Wait()
}

//// 实现自旋锁
//var Count int32 = 1
//
//func SpinLock(i int32, fn func()) {
//	for {
//		if n := atomic.LoadInt32(&Count); n == i {
//			fn()
//			atomic.AddInt32(&Count, 1)
//			break
//		}
//		time.Sleep(time.Nanosecond)
//	}
//}
//func main() {
//	for i := int32(1); i < 4; i++ {
//		go func(i int32) {
//			fn := func() {
//				fmt.Println(i)
//			}
//			SpinLock(i, fn)
//		}(i)
//	}
//	SpinLock(4, func() {})
//}
