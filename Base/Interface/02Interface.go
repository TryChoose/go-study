package main

import "fmt"

//汽车接口
type move interface {
	run()
}
type bicycle struct {
	brand string
	wheel int
}
type car struct {
	brand string
	wheel int
}

//该函数接受一个move类型的变量 接口变量
func moveChe(m move) {
	m.run()
}
func (b bicycle) run() {
	fmt.Printf("我是%s,我有%d轮子\n", b.brand, b.wheel)
}
func (c car) run() {
	fmt.Printf("我是%s,我有%d轮子\n", c.brand, c.wheel)
}

func main() {
	b := bicycle{
		brand: "自行车",
		wheel: 2,
	}
	c := car{
		brand: "机动车",
		wheel: 4,
	}
	moveChe(b)
	moveChe(c)
}
