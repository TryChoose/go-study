package main

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
// 两个channel  一个控制字符 一个控制数字
/*
 Done后面就不能使用Add
*/
func main() {
	l, n := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-n:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				l <- true
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-l:
				if i >= 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				n <- true
			}
		}
	}(&wait)
	n <- true
	wait.Wait()
}
