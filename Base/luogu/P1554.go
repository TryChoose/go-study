package main

import "fmt"

func main() {
	m1 := make(map[int]int)
	var m, n int
	fmt.Scan(&m, &n)
	for i := 0; i <= 9; i++ {
		m1[i] = 0
	}
	t := 0
	for m <= n {
		t = m
		for t != 0 {
			if _, ok := m1[t%10]; !ok {
				m1[t%10] = 1
			} else {
				m1[t%10] += 1
			}
			t = t / 10
		}
		m++
	}
	a := []int{
		0: m1[0],
		1: m1[1],
		2: m1[2],
		3: m1[3],
		4: m1[4],
		5: m1[5],
		6: m1[6],
		7: m1[7],
		8: m1[8],
		9: m1[9]}
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}
