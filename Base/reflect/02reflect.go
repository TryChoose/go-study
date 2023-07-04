package main

import (
	"fmt"
	"reflect"
)

type student struct {
	ID    int64
	Name  string
	Class string
}

func newStudent(ID int64, name string, class string) *student {
	return &student{ID: ID, Name: name, Class: class}
}

type teacher struct {
	ID    int
	Name  string
	Class string
	Age   string
}

func newTeacher(ID int, name string, class string, age string) *teacher {
	return &teacher{ID: ID, Name: name, Class: class, Age: age}
}

/*
使用反射更改结构体里面 变量的值
1.判断是否是指针
2.在判断指针里面包含的是否是结构体指针
3.
*/
func reflect01(a *student) {
	v := reflect.ValueOf(a).Elem()
	//获取指针指向的元素
	////取字段
	//field := v.Field(0)
	//fmt.Println(field) //101
	name := v.FieldByName("ID")
	fmt.Println(name.Kind())
	if name.Kind() == reflect.Int64 {
		name.SetInt(int64(102))
	}

}
func main() {
	s1 := newStudent(101, "张三", "一班")
	fmt.Println(s1.ID, s1.Name, s1.Class)
	reflect01(s1)
	fmt.Println(s1.ID, s1.Name, s1.Class)
}
