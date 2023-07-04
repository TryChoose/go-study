package main

import (
	"fmt"
	"sync"
)

func f3(ch chan int) {
	//for {
	//	v, ok := <-ch
	//	if !ok {
	//		fmt.Println("通道已关闭")
	//		break
	//	}
	//	fmt.Printf("v:%#v ok:%#v\n", v, ok)
	//}
}

// demo1 通道误用导致的bug
func demo1() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				//if !ok{
				//	break
				//}
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func main() {
	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//close(ch)
	//f3(ch)
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	demo1()
}
