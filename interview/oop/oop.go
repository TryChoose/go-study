package main

import "fmt"

// go语言实现面向对象
// 面向对象实现
// 封装和继承是由 结构体实现 多态由接口实现
type A interface {
	setName(name string)
	eat()
	sleep()
}
type Animal struct {
	name string
}
type Cat struct {
	Animal
	age int
}
type Dog struct {
	Animal
	weight int
}

func NewAnimal() *Animal {
	return &Animal{}
}
func (a *Animal) setName(name string) {
	a.name = name
}
func (a *Animal) eat() {
	fmt.Println("I am", a.name+",I am eating")
}
func (a *Animal) sleep() {
	fmt.Println("I am", a.name+",I am sleeping")

}
func (c *Cat) setAge() {
	c.age = 10
}
func (d *Dog) setWeight() {
	d.weight = 100
}
func come(a A) {
	a.sleep()
}
func main() {
	//dog.eat()
	//dog.sleep()
	//come(dog)
	dog := Dog{Animal: *NewAnimal()}
	dog.setName("dog")
	dog.sleep()
	dog.eat()
	dog.setWeight()
	fmt.Println(dog.weight)
	cat := Cat{Animal: *NewAnimal()}
	cat.setName("cat")
	cat.sleep()
	cat.eat()
	cat.setAge()
	fmt.Println(cat.age)
	come(&dog)
	come(&cat)
}
