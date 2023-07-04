package main

import "fmt"

//构造函数

type student struct {
	name   string
	age    int
	gender string
}

//构造函数new
//返回值的是结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针,减少内存的消耗
func newStudent(name string, age int, gender string) student {
	return student{
		name:   name,
		age:    age,
		gender: gender,
	}
}
func newStudent1(name string, age int, gender string) *student {
	return &student{
		name:   name,
		age:    age,
		gender: gender,
	}
}
func main() {
	p1 := newStudent("小明", 20, "男")
	fmt.Println(p1)
	p2 := newStudent1("张三", 29, "女") //返回的是结构体的内存地址
	fmt.Printf("%#v\n", p2.name)     //取出结构体里面的值
	fmt.Printf("%#v", p2)            //&main.student{name:"张三", age:29, gender:"女"
}
