package main

import (
	"fmt"
	"time"
)

func f(i int) {
	fmt.Println("Hello,World", i)
}
func main() {
	//for i := 0; i <100 ; i++ {
	//	go f(i)
	//}
	fmt.Println("main")
	//time.Sleep(time.Second)
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println("Hello,World", i)
		}
	}()
	//fmt.Println("main")
	time.Sleep(time.Second)
}
