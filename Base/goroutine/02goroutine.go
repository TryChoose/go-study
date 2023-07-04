package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func randDemo() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}

}

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	//time.Sleep(time.Second*time.Duration(rand.Intn(3)))
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	defer wg.Done()
	//time.Sleep(time.Second*time.Duration(rand.Intn(3)))
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	//randDemo()
	runtime.GOMAXPROCS(1)
	//numCPU := runtime.NumCPU()
	//fmt.Println(numCPU)
	wg.Add(2)
	go a()
	go b()

	wg.Wait()
}
