package main

import "fmt"

//空接口
func main() {
	m1 := make(map[string]interface{}, 10)
	m1["张三"] = 12
	m1["AVD"] = "ASD"
	m1["DS"] = [...]string{"ASD", "SSD", "DSD"}
	fmt.Printf("%#v", m1)
}
