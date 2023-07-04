package main

import "fmt"

//指针
func main() {
	//1.& 取地址
	//n:=18
	//p:=&n
	////取值
	//fmt.Println(n,*p)
	//fmt.Printf("%T",p)

	//var a *int //invalid memory address or nil pointer dereference
	//*a=100
	//fmt.Println(*a)
	//new 函数申请一个内存地址
	var a = new(int)
	fmt.Println(a)
}
