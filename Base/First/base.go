package main

import "fmt"

func main() {
	//// %d 表示整型数字，%s 表示字符串
	//var stockcode=123
	//var enddate="2020-12-31"
	//var url="Code=%d&endDate=%s"
	//var target_url=fmt.Sprintf(url,stockcode,enddate)
	//fmt.Println(target_url)

	/**
		声明变量
		 第一种方式 var方式：
		 var 变量名 类型
		 var 变量名,变量名,.....变量名 类型

		 var 变量名=value  不能换行赋值
		 第二种 :=声明
		  var intval int
		  intval=1
	      相等于 intval:=1

		const 定义常量
	*/
	//var a ,b  int32
	//var c string="name"
	//a=1
	//b=2
	//var d=20000000
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//var e int
	//e=0
	//fmt.Println(e)
	//f:=1
	//fmt.Println(f)
	//g:=true
	//fmt.Println(g)

	/**
	const 定义常量
	*/
	//const LENGTH int = 10
	//const WIDTH int = 5
	//var area int
	//const a, b, c = 1, false, "str" //多重赋值
	//
	//area = LENGTH * WIDTH
	//
	//fmt.Printf("面积为 : %d", area)
	//println()
	//println(a, b, c)
	//const (
	//	a="abc"
	//	b=len(a)
	//	c=unsafe.Sizeof(a)
	//)

	//*
	//iota 特殊常量
	//const (
	//	a=iota  //0
	//	b		//1
	//	c        //2
	//	d="ha"   //值为 ha iota=3
	//	e      //ha  4
	//	f=iota  //5
	//	g  //6
	//	i //7
	//)
	//fmt.Println(a,b,c,d,e,f,g,i)

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c) //31
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c) //11
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c) //210
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c) //2
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c) //1
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a) //22
	a = 21                            // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a) //20
}
