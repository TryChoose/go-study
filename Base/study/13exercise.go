package main

import (
	"fmt"
	"strings"
)

//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func main() {
	//s := [...]string{
	//	"how",
	//	"do",
	//	"you",
	//	"do",
	//}
	//m1 := make(map[string]int, 10)
	//m1["how"] = 0
	//m1["do"] = 0
	//m1["you"] = 0
	//for _, s2 := range s {
	//	//fmt.Println(s2)
	//	if s2 == "how" {
	//		m1["how"]++
	//	}
	//	if s2 == "do" {
	//		m1["do"]++
	//	}
	//	if s2 == "you" {
	//		m1["you"]++
	//	}
	//}
	//
	//for k, v := range m1 {
	//	fmt.Println(k, v)
	//}
	//s := "Hello,how do you do"
	//sl := make([]string, 0, 6) //切片
	////word := ""
	//var word string =""
	//for _, w := range s { // 遍历字符串
	//	if !unicode.IsLetter(w) { // 遇到非字母元素，说明一个单词结束，将单词存入slice中，并且重置word变量
	//		if word != ""{ // 由于可能存在两个或者多个的非字母字符相邻，所以在录入slice之前应该进行一次判断
	//			sl = append(sl, word)
	//		}
	//		word = ""
	//		continue
	//	} else {
	//		word = fmt.Sprintf("%s%c", word, w) // 字母元素拼接到word后面，组成单词
	//	}
	//}
	//fmt.Println(sl)
	//wordMap := make(map[string]int, 6) //map
	//for _, k := range sl {
	//	wordMap[k]++
	//}
	//for k, v := range wordMap {
	//	fmt.Println(k, v)
	//}
	words := "how do you do"
	wordSlice := strings.Split(words, " ")
	wordMap := make(map[string]int, 10)
	for _, s := range wordSlice {
		_, ok := wordMap[s]
		if ok {
			wordMap[s]++
		} else {
			wordMap[s] = 1
		}
	}
	for k, v := range wordMap {
		fmt.Printf("单词:%s 个数:%d\n", k, v)
	}
	//fmt.Println(wordMap)
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])

}
