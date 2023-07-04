package main

import "fmt"

func main() {
	var grade string = "B"
	switch {
	case grade == "B":
		fmt.Printf("良好\n")
	}
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("x的类型：%T", i)
	case int:
		fmt.Println("x是int 型")
	case bool, string:
		fmt.Println("x是bool或string型")
	case func(int) int: //函数类型变量 func(变量名 变量类型)返回值
		fmt.Println("x是func(int)型")
	}
	var c1, c2, c3 chan int //通道类型
	var i2 int
	select { //select 随机执行case
	case i1 := <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}
