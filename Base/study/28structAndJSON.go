package main

import (
	"encoding/json"
	"fmt"
)

//结构体与JSON

//将结构体转为JSON文件 -->序列化
//将JSON文件转为结构体 -->反序列化
//注意事项:结构体里面的变量名需要大写
type t struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	t1 := t{
		Name: "GOlang",
		Age:  19,
	}
	//序列化
	b, err := json.Marshal(t1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", string(b))
	//反序列化
	str := `{"name":"张三","age":10}`
	var t2 t
	json.Unmarshal([]byte(str), &t2)
	fmt.Printf("%v", t2)
}
