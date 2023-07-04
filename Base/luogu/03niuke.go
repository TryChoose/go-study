package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	nanoTime := int64(time.Now().Nanosecond())
	rand.Seed(nanoTime)
	k := rand.Int63n(int64(a[n-1]))
	count := 0
	for _, v := range a {
		b := math.Abs(float64(k - int64(v)))
		if b > float64(x) {
			count++
		}
	}
	fmt.Println(count)
}
