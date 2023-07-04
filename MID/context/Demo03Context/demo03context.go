package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
使用传统方式关闭子 goroutine
*/
var wg sync.WaitGroup
var lock sync.Mutex
var i int = 0
var j int = 1000

func f2(ctx context.Context) {
	defer wg.Done()
FLoop:
	for {
		fmt.Println(j)
		lock.Lock()
		j++
		lock.Unlock()
		time.Sleep(time.Millisecond * 100)
		select {
		case <-ctx.Done():
			break FLoop
		default:
		}
	}
}
func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FLoop:
	for {
		fmt.Println(i)
		lock.Lock()
		i++
		lock.Unlock()
		time.Sleep(time.Millisecond * 100)
		select {
		case <-ctx.Done():
			break FLoop
		default:
		}
	}
}
func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 10)
	cancelFunc()
	wg.Wait()
}
