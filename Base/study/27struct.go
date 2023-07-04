package main

import "fmt"

//模拟面向对象的继承

type person struct {
	idCard int
}

type students struct {
	name string
	person
}

func newStudents(name string, person person) *students {
	return &students{name: name, person: person}
}
func newPerson(idCard int) *person {
	return &person{
		idCard: idCard,
	}
}
func (p person) read() {
	fmt.Println(p.idCard)
}
func (s students) look() {
	fmt.Println(s.name)
}
func main() {
	p := person{01}
	s1 := newStudents("张三", p)
	s1.look()
	s1.read()
}
