package main

import "fmt"

//嵌套结构体
type address struct {
	province string
	city     string
}
type p struct {
	name    string
	age     int
	address address
}

type p1 struct {
	name    string
	address address
}

func main() {
	p1 := p{
		name: "张三",
		age:  20,
		address: address{
			province: "湖北",
			city:     "武汉",
		},
	}
	fmt.Println(p1, p1.address.province)
}
