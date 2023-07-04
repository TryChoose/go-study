package main

import (
	"bytes"
	"fmt"
)

//切片的比较
//对于字节[]byte 切片我们可以采用 bytes.Equal
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func main() {
	//字节切片对比
	a := []byte{'1', '2'}
	b := []byte{'3', '4'}
	fmt.Println(bytes.Equal(a, b))
	c := []string{"张三李四"}
	d := []string{"张三李四"}
	fmt.Println(equal(c, d))

}
