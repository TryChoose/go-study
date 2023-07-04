package main

import "fmt"

func main() {
	//普通形式
	//for i := 0; i <10 ; i++ {
	//	fmt.Println(i)
	//}
	//变种1
	//i:=4
	//for;i<10;i++{
	//	fmt.Println(i)
	//}
	//变种2
	//i:=5
	//for i<10 {
	//	fmt.Println(i)
	//	i++
	//}
	////死循环
	//for {
	//	fmt.Println("123")
	//}
	//for range循环
	s := "HelloWorld做广告"
	//for i,v:=range s{
	//	fmt.Printf("%d %c\n",i,v)
	//}
	/*for v:=range s{
		fmt.Printf("%d\n",v)
	}*/
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
	//for i := 1; i <=9 ; i++ {
	//	for j := 1; j <=i; j++ {
	//		fmt.Printf("%d*%d=%d\t",i,j,i*j)
	//	}
	//	fmt.Println()
	//}
}
