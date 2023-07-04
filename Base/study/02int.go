package main

import "fmt"

func main() {
	var i1 = 10
	var i2 = 077
	var i3 = 0xABDC
	i4 := int8(10)
	fmt.Printf("%d\n", i1)
	fmt.Printf("%d\n", i2)
	fmt.Printf("%d\n", i3)
	fmt.Printf("%d\n", i4)
	fmt.Printf("%T\n", i4)

}
