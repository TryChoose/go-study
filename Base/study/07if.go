package main

import "fmt"

func main() {
	//普通if语句形式
	//age:=20
	//if age>=18 {
	//	fmt.Printf("你已经成年了！\n")
	//}else {
	//	fmt.Println("你还未成年！\n" )
	//}
	//fmt.Println(age)
	//特殊形式if
	if age := 20; age >= 35 {
		fmt.Printf("中年！\n")
	} else if age >= 18 {
		fmt.Println("青年\n")
	} else {
		fmt.Println("未成年!\n")
	}
	/*
		总结：两种形式的if语句，区别作用域不同,普通形式的if语句的变量作用域在整个函数中
		特殊形式的作用域在于if语句之内，外部无法引用到该变量
	*/
}
