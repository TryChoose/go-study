package main

import "fmt"

//类型断言
func assert1(a interface{}) {
	s, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("猜对了！%T,%v", s, s)
	}
}
func assert2(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Println(t)
	case int:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	case float64:
		fmt.Println(t)
	}
}
func main() {
	//assert1("100")
	assert2("12")
}
