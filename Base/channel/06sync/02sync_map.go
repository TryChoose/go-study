package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go内置的map不是并发安全的
var m1 = make(map[string]int)
var wg sync.WaitGroup
var m2 sync.Map
var lock sync.Mutex

func set(key string, value int) {
	m1[key] = value
}
func get(key string) int {
	return m1[key]
}

//	func main() {
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go func(i int) {
//				key := strconv.Itoa(i)
//				lock.Lock()
//				set(key, i)
//				lock.Unlock()
//				fmt.Println(key, get(key))
//				wg.Done()
//			}(i)
//		}
//		wg.Wait()
//	}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			m2.Store(key, i)
			value, _ := m2.Load(key)
			fmt.Printf("%T:%+v %T:%+v\n", key, value, key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
