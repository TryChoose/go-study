package main

import "fmt"

//方法和接收者

type dog struct {
	name string
	age  int
}

//newDog构造函数
func newDog(name string, age int) *dog {
	return &dog{name: name, age: age}
}

//值类型接收者
func (d dog) eat() {
	fmt.Printf("%s:吃好喝好...\n", d.name)
}

//指针类型接收者
func (d *dog) sleep() {
	d.name = "张三"
	fmt.Printf("%s:睡好觉...", d.name)
}
func main() {
	d1 := newDog("卡里", 20)
	d1.eat()
	d1.sleep()
}
