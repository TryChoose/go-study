package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序.
*/
var wg sync.WaitGroup
var ch1 chan int64
var ch2 chan int64

/*
分析:
1.首先第一步 需要将一个随机数取出来， 然后放入到通道里面 （接收通道）
2.将这个通道里面的值取出来放到另外一个接收通道里面
3.遍历接受通道ok
*/
//
type t1 struct {
	value int64
}
type res struct {
	value *t1
	sum   int64
}

func f1(ch1 chan<- *t1) {
	defer wg.Done()
	//定义随机数
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Int63()
		t := &t1{
			value: x,
		}
		ch1 <- t
		time.Sleep(time.Millisecond * 200)
	}
}
func f2(ch1 <-chan *t1, ch2 chan<- *res) {
	defer wg.Done()
	for {
		var sum int64 = 0
		a := <-ch1 //接收的是一个结构体t1
		n := a.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		r := &res{
			value: a,
			sum:   sum,
		}
		ch2 <- r
	}
}
func main() {
	ch1 := make(chan *t1, 100)
	ch2 := make(chan *res, 100)
	wg.Add(1)
	go f1(ch1)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go f2(ch1, ch2)
	}
	for i := range ch2 {
		fmt.Println("value: ", i.value.value, "sum:", i.sum)
	}
	wg.Wait()
}
