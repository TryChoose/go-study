package main

import "fmt"

//面试题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

//最后打印:
/*
 A  1 2 3
 B  10 2 12
 BB 10,12 22
 AA 1 3  4
*/
// 1. x赋值为1 ,y赋值为2
// 2.defer calc("AA", x, calc("A", x, y))
// 3.calc("A", x, y)  "A" 1 2 3
// 4.defer calc("AA", x,3) 压栈  x赋值为1
// 5.x=10
// 6.defer calc("BB", x, calc("B", x, y))
// 7.calc("B", x, y) "B",1O,2 12
// 8.defer calc("BB", x, 12) 压栈  x=10
// 9.y=20
// 10.出栈 “BB” 10,12 22
// 11.出栈 “AA”  2 3  5
