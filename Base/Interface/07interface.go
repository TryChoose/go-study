package main

import (
	"fmt"
)

type animaler interface {
	sleep()
	eat()
}
type tiger struct {
}

func (t *tiger) sleep() {
	fmt.Println("tiger sleep")
}
func (t *tiger) eat() {
	fmt.Println("tiger eat")
}
func come(a animaler) {
	a.sleep()
	a.eat()
}
func main() {
	come(&tiger{})
}
