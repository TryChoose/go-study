package main

//import "fmt"
//
//////defer语句
////func f1()  {
////	fmt.Println("start")
////	defer fmt.Println("*******")
////	defer fmt.Println("-------")
////	defer fmt.Println("_______")
////	fmt.Println("end")
////}
////func main() {
////	f1()
////}
//
//func f1() int {
//	/*
//	执行顺序：
//	1.赋值 已经为5
//	2.defer 改变为6
//	3.return 最后返回赋值的值 ,并不返回6
//	 */
//	x := 5
//	defer func() {
//		x++
//		fmt.Println(x)
//	}()
//	return x
//}
//
//func f2() (x int) {
//	/*
//	执行顺序：
//	1.赋值 x=5
//	2.defer x++ 6
//	3.返回 6
//	 */
//	defer func() {
//		x++
//		fmt.Println(x)
//	}()
//	return 5  //相当于给赋值 x=5
//}
//
//func f3() (y int) {
//	/*
//		执行顺序：
//		1.赋值 y=x=5
//		2.defer x++ 6
//		3.返回 5
//	*/
//	x := 5
//	defer func() {
//		x++
//		fmt.Println(x)
//
//	}()
//	return x  //y=x=5
//}
//func f4() (x int) {
//	/*
//		执行顺序：
//		1.赋值 x=5
//		2.defer x++ 1
//		3.返回 5
//	*/
//	defer func(x int) {
//		x++
//		fmt.Println(x)
//	}(x)
//	return 5
//}
//func main() {
//	//fmt.Println(f1())
//	//fmt.Println(f2())
//	//fmt.Println(f3())
//	//fmt.Println(f4())
//}
func main() {

}
