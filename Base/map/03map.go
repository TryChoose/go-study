package main

import "fmt"

//两个map进行比较

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
func main() {
	m1 := map[string]int{"张三": 1, "李四": 2}
	m2 := map[string]int{"李四": 2, "张三": 1}
	fmt.Println(equal(m1, m2))
}
