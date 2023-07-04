package main

import (
	"fmt"
	"os"
)

//命令行参数
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	for _, v := range os.Args {
		fmt.Println(v)
	}

	fmt.Println(s)
}
