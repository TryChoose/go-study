package main

import (
	"fmt"
	"strconv"
)

//import "fmt"
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	// write code here
	str := strconv.Itoa(x)
	k := len(str)
	j := k - 1
	for i := 0; i < k; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
	//var a = make([]int, 0, 10)
	//var b = make([]int, 0, 10)
	//var flag bool
	//n, i:= 0, 0
	//var m int
	//for {
	//	if x <= 0 {
	//		break
	//	} else {
	//
	//		m = x % 10
	//		x /= 10
	//		a=append(a,m)
	//		n++
	//		i++
	//	}
	//}
	//fmt.Println(a,b)
	////for j = 0; j < n; j++ {
	////	b=append(a,)
	////}
	//b=append(b,a[0:]...)
	//fmt.Println(a,b)
	//for i = 0; i < n; i++ {
	//	if b[i] != a[i] {
	//		flag=false
	//	}else {
	//		flag=true
	//	}
	//}
	//return  flag
}
func main() {
	fmt.Println(isPalindrome(991))

}
