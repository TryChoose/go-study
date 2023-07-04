package main

import (
	"fmt"
	"time"
)

func main() {
	strChan := make(chan int, 1)
	numChan := make(chan int, 1)
	strChan <- 0
	go func() {
		for i := 65; i <= 90; i++ {
			<-strChan
			fmt.Printf("%v ", string(rune(i)))
			numChan <- i
		}
		return
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Printf("%v ", i)
			strChan <- i
		}
		return
	}()
	time.Sleep(1 * time.Second)
	fmt.Println()
}
