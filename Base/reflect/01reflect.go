package main

import (
	"fmt"
	"reflect"
)

type cat struct {
}

//反射
func reflectType(x interface{}) {
	v := reflect.TypeOf(x) //具体的类型
	vkind := v.Kind()
	fmt.Printf("type:%v  kind:%v\n", v, vkind)
}

func main() {
	var a float64 = 3.14
	reflectType(a) //type:float64  kind:float64
	var b int64 = 100
	reflectType(b) //type:int64  kind:int64
	var c = cat{}
	reflectType(c) //type:main.cat  kind:struct
}
