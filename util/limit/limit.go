package main

import (
	"fmt"
	"time"
)

var limitCh = make(chan struct{}, 10)

// limit 限制并发数
func limit() {
	limitCh <- struct{}{}
	time.Sleep(time.Millisecond * 100)
	<-limitCh
}
func main() {
	for i := 0; i < 100; i++ {
		go func() {
			limit()
			fmt.Println("业务逻辑")
		}()
	}
	time.Sleep(time.Second * 10)
}
