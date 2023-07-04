package main

import "fmt"

func canConstruct(des string, src string) bool {
	// write code here
	s1 := []byte(src)
	m1 := make(map[byte]int)
	for _, v := range s1 {
		m1[v]++
	}
	for i, i2 := range m1 {
		fmt.Printf("%c,%d\n", i, i2)
	}
	s2 := []byte(des)
	for _, v := range s2 {
		if m1[v] == 0 {
			return false
		}
		m1[v]--
	}
	return true
}
func main() {
	s := canConstruct("ab", "aab")
	fmt.Println(s)
}
