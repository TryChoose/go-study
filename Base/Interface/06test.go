package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

//指针接收者
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
