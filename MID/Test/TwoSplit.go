package Test

import (
	"fmt"
)

func Solution(str string) []string {
	n := len(str)
	if n == 1 {
		return []string{str}
	}
	res := make([]string, (len(str)/2)+1)
	s := ""
	j := 0
	for i := 1; i < len(str); i++ {
		s += string(str[i-1])
		fmt.Println(s)
		if len(s)%2 == 0 {
			res[j] = s
			s = ""
			j++
		}
	}
	s += string(str[n-1]) + "_"
	res[j] = s
	return res
}

//func CreatePhoneNumber(numbers [10]uint) string {
//	bs := []byte{}
//	//rand.Seed(time.Millisecond.M)
//}
