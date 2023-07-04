package main

import "fmt"

//实现二维数组
type twoGroup interface {
	Init(mX, mY, r int)
	Set(x, y, r int)
	Get(x, y int)
}

type you struct {
	data []int
	maxX int
	maxY int
}

func newYou(data []int) *you {
	return &you{data: data}
}

//初始化方法
func (a *you) Init(x, y, r int) {
	a.maxX = x
	a.maxY = y
	a.data = make([]int, a.maxX*a.maxY)
	for i := 0; i < a.maxX*a.maxY; i++ {
		a.data[i] = r
	}
}
func (a *you) Set(x, y, r int) {
	var n int
	n = x*a.maxY + y
	a.data[n] = r
}

func (a *you) Get(x, y int) int {
	var n int
	n = x*a.maxY + y
	return a.data[n]
}

func main() {
	a := []int{}
	yObj := newYou(a)
	yObj.Init(2, 4, 0)
	fmt.Println(yObj.maxX, yObj.maxY)
	yObj.Set(1, 1, 100)
	get := yObj.Get(1, 1)
	fmt.Println(get)
}
