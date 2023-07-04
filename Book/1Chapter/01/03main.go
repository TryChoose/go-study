package main

import (
	"bufio"
	"fmt"
	"os"
)

//找出重复行
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "0" {
			break
		}
	}
	for l, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, l)
		}
	}

}
