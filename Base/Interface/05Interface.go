package main

//一个类型是实现多个接口
//嵌套接口
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat()
}
type cat struct {
}

func (c cat) move() {

}
func (c cat) eat() {

}
func main() {

}
