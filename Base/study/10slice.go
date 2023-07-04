package main

import (
	"fmt"
)

//切片slice
func main() {
	//切片的定义
	//var 变量名 []T
	//var a1 [] int
	//var a2 []string
	////初始化
	//a1=[]int{1,2,3,4}
	//a2=[]string{"张三","李四","王五"}
	//fmt.Println(a1,a2)

	////长度和容量
	//fmt.Printf("len(a1):%d,cap(a1)%d\n",len(a1),cap(a1))
	//fmt.Printf("len(a2):%d,cap(a2)%d\n",len(a2),cap(a2))

	////由数组得到切片
	//a3:=[...]int{1,2,3,4,5,6,7,8,9,10}
	//a4:=a3[0:5]
	//fmt.Println(a3,a4)
	//a5:=a3[:3]
	//a6:=a3[2:]
	//a7:=a3[:]
	//fmt.Println(a5,a6,a7)
	//make()构造切片

	//a := make([]string, 10, 20)
	//fmt.Printf("%v,%d,%d,%T\n", a, len(a), cap(a), a)
	//
	////切片的赋值
	//a1 := []int{1, 2, 3, 5}
	//a2 := a1
	//fmt.Println(a1, a2)
	//a1[0] = 100
	//fmt.Println(a1, a2)
	//for i := 0; i < len(a2); i++ {
	//	fmt.Println(i, a2[i])
	//}
	////for range 遍历
	//for i, v := range a1 {
	//	fmt.Println(i, v)
	//}

	//append()为切片追加元素
	//a1:=[]string{"北京","上海","武汉"}
	//fmt.Println(a1)
	//fmt.Println(len(a1),cap(a1))
	////调用append()必须用原来的切片变量接受返回值
	////append追加元素,原来的底层数组放不下时,Go底层就会把底层数组换一个
	////必须使用变量接受append的返回值
	//a1=append(a1,"长沙")
	//fmt.Println(a1)
	//fmt.Println(len(a1),cap(a1))
	//a1=append(a1,"重庆","广州")
	//fmt.Println(a1)
	//fmt.Println(len(a1),cap(a1))
	//a1=append(a1,"厦门")
	//fmt.Println(a1)
	//fmt.Println(len(a1),cap(a1))
	//a2:=[]string{"西安","苏州","杭州"}
	//a1=append(a1,a2...) //...表示拆开,意思就是说 将a2切片拆开一个个放入a1切片里面
	//fmt.Println(a1)
	//fmt.Println(len(a1),cap(a1))

	//copy()
	//a1:=[]int{1,2,3,5}
	//a2:=a1;
	//fmt.Println(a1,a2)
	//var a3=make([]int,4,6)
	//copy(a3,a1)
	//fmt.Println(a1,a2,a3)
	//a1[0]=100
	//a1:=[]int{1,2,3,4,5,6,7,8,9,10}
	//a1=append(a1[:2],a1[3:]...)
	//fmt.Println(a1)

	//a:=[...]int{1,3,5,6,7,8}
	//var a1 []int
	//fmt.Println(a)
	//a1=a[1:4]
	//fmt.Println(a1)
	//a1=append(a[:2],a[3:]...)
	//fmt.Println(a1)
	//fmt.Println(a)

	var a = make([]int, 0, 10) //开辟
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	//fmt.Println(cap(a))
	//
	//a1:=[...]int{1,5,7,2,12,24} //数组
	//sort.Ints(a1[:])//切片
	//fmt.Println(a1)

	//切片删除元素 或者说 移动元素
	//a1:=[...]int{1,5,7,3,61,37,123,52,51}
	//a2:=a1[:]
	//a2=append(a1[:4],a1[6:]...)
	//fmt.Println(a2)
	//fmt.Println(a1)

}
