package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//
// f1
//  @Description: 主goroutine使用全局变量通知子goroutine退出
//

var wg sync.WaitGroup
var flag bool

func f1() {
	defer wg.Done()
	for {
		fmt.Println("张三！！！！！")
		if flag {
			break
		}
	}

}

//
// f2
//  @Description: 使用channel+select 退出
//
var fc = make(chan bool, 1)

func f2() {
	defer wg.Done()
	for {
		fmt.Println("张三！！！！！")
		select {
		case <-fc:
			break
		default:
		}
	}
}

//
// f3
//  @Description: 使用context
//
func f3(c context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("张三！！！！！")
		select {
		case <-c.Done():
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	//go f1()
	go f3(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	//fc <- true
	//fmt.Println("f1 退出")
	//fmt.Println("f2 退出")
	wg.Wait()
}
