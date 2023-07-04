package main

import (
	"fmt"
	"path"
	"runtime"
)

//调用栈
func f2() {
	f1()
}
func f1() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(path.Base(file)) //01runtime_demo.go
	fmt.Println(funcName)
	fmt.Println(file, line, ok)
}
func main() {
	//f1()
	f2()
}
