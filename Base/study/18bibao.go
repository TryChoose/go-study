package main

//
//import (
//	"fmt"
//	"strings"
//)
//
////闭包的含义:
//	//闭包是一个函数,这个函数包含了外部作用域的一个变量
////func f1(x int) func(int) int{
////	return func(y int) int {
////		x+=y
////		return x
////	}
////}
////func f1(f func()){
////	fmt.Println("this is f1")
////	f()
////}
////func f2(x,y int){
////	fmt.Println("this is f2")
////	fmt.Println(x+y)
////	//f1(func() {
////	//	fmt.Println(x+y)
////	//})
////}
////func f3(x,y int)(func ()){
////	temp:= func() {
////		fmt.Println(x+y)
////	}
////	return  temp
////}
//func makeSuffixFunc(suffix string) func(string) string{
//	return func(name string) string {
//		if !strings.HasSuffix(name,suffix) {
//			return  name+suffix
//		}
//		return  name
//	}
//}
//func main() {
//	//res:=f1(100)  //通过调用f1返回的返回值是一个 带一个int类型参数 返回值为int的函数类型
//	//res2:=res(200)//调用res返回值是一个int类型
//	//fmt.Println(res2 )
//	//res := f3(10, 20)
//	//f1(res)
//	jpgFunc:=makeSuffixFunc(".jpg")
//	txtFunc:=makeSuffixFunc(".txt")
//	fmt.Println(jpgFunc("test"))
//	fmt.Println(txtFunc("test"))
//}
//
//
func main() {

}
