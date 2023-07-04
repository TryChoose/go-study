package main

import "fmt"

//接口里面只能有方法
//面向接口编程
type student interface {
	introduce()
	read()
	write()
	eat()
	move()
}
type man struct {
	sno       int
	name      string
	gender    string
	readBook  string
	writeWork string
	eatFood   string
	moveStep  int
}
type woman struct {
	sno       int
	name      string
	gender    string
	readBook  string
	writeWork string
	eatFood   string
	moveStep  int
}

func newMan(sno int, name string, gender string, readBook string, writeWork string, eatFood string, moveStep int) *man {
	return &man{sno: sno, name: name, gender: gender, readBook: readBook, writeWork: writeWork, eatFood: eatFood, moveStep: moveStep}
}

func newWoman(sno int, name string, gender string, readBook string, writeWork string, eatFood string, moveStep int) *woman {
	return &woman{sno: sno, name: name, gender: gender, readBook: readBook, writeWork: writeWork, eatFood: eatFood, moveStep: moveStep}
}

func (m man) introduce() {
	fmt.Printf("大家好！我是%s,我的学号是%d号,我是一名%s生.\n", m.name, m.sno, m.gender)
}
func (w woman) introduce() {
	fmt.Printf("大家好,我是%s,我的学号是%d号,我是一名%s生.\n", w.name, w.sno, w.gender)
}
func (m man) read() {
	fmt.Printf("我喜欢看%s\n", m.readBook)
}
func (w woman) read() {
	fmt.Printf("我喜欢看%s\n", w.readBook)
}
func (m man) write() {
	fmt.Printf("我最爱写%s\n", m.writeWork)
}
func (w woman) write() {
	fmt.Printf("我最爱写%s\n", w.writeWork)
}
func (m man) eat() {
	fmt.Printf("我喜欢%s\n", m.eatFood)
}
func (w woman) eat() {
	fmt.Printf("我喜欢%s\n", w.eatFood)
}
func (m man) move() {
	fmt.Printf("我天天走%d步\n", m.moveStep)
}
func (w woman) move() {
	fmt.Printf("我天天走%d步\n", w.moveStep)
}
func stu(s student) {
	s.introduce()
	s.read()
	s.write()
	s.eat()
	s.move()
}
func main() {
	m1 := newMan(1012, "小张", "男", "《铁臂阿童木》", "《五年高考 三年模拟》", "吃好吃的,嘎嘎吃", 10000)
	m2 := newWoman(1012, "小李", "女", "《芭比娃娃》", "《霸道总裁爱上我》", "吃美食，嘎嘎吃", 10000)
	stu(m1)
	stu(m2)
}
