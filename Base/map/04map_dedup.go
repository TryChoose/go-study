package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
		if line == " " {
			break
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "04map_dedup:%v\n", err)
		os.Exit(1)
	}
}
