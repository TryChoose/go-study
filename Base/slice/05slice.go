package main

import "fmt"

//去除相邻的重复字符串元素

func loadRemove(sliceString []string) []string {
	//1.直接使用map存放结果 每个字符只需要一次
	m1 := make(map[string]bool)
	res := []string{}
	for _, v := range sliceString {
		if _, ok := m1[v]; !ok {
			res = append(res, v)
			m1[v] = true
		}
	}
	return res
}
func main() {

	s := []string{"张三", "李四", "王五", "肖战", "王五"}
	fmt.Println(loadRemove(s))

}
