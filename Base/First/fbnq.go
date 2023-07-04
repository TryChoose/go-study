package main

import "fmt"

func fun(n int) {
	a := [5]int{1, 1}
	//var i int
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
		fmt.Println(a[i])
	}

}
func main() {
	//fun(5)
}
