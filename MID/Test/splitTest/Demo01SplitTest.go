package splitTest

import "strings"

// 单元测试小案例
// 切割字符串测试
func Split(str string, sep string) []string {
	ret := make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep) //求出字串在原字符串中的索引
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}

// fib.go

// Fib 是一个计算第n个斐波那契数的函数
//func Fib(n int) int {
//	if n < 2 {
//		return n
//	}
//	return Fib(n-1) + Fib(n-2)
//}
