package main

import "fmt"

// panic 和 recover
//panic 崩溃程序
//recover 获取函数 必须和defer一起使用

func funA() {
	fmt.Println("A")
}
func funB() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("error！！")
	fmt.Println("B")
}
func funC() {
	fmt.Println("C")
}

func main() {
	funA()
	funB()
	funC()
}
