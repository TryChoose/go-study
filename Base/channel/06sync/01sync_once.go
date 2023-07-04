package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	once.Do(func() {
		fmt.Println("Only once")
	})

}
