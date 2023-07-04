package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
随机生成8-16位密码
包含：大小写,数字，@!.
*/
const (
	u = "WERTYUIOPASDFGHJKLZXCVBNM"
	l = "qwertyuiopasdfghjklzxcvbnm"
	n = "1234567890"
	t = "!@."
)

func getRandomWithAll(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}

func isValid(str string) bool {
	//判断是否是有效的密码，密码必须以大写开头 包括大小写数字
	b := []byte(str)
	if !(b[0] >= 'A' && b[0] <= 'Z') {
		return false
	}
	if !(strings.ContainsAny(str, u) && strings.ContainsAny(str, l) && strings.ContainsAny(str, n) && strings.ContainsAny(str, t)) {
		return false
	}
	return true
}
func producePassword(n int64) string {
	//随机数种子
	var letters = []byte("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890!@.")
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, n)
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	if isValid(string(res)) {
		return string(res)
	} else {
		n1 := getRandomWithAll(8, 16)
		password := producePassword(n1)
		valid := isValid(password)
		if valid {
			return password
		}
	}
	return ""
}
func main() {
	n := getRandomWithAll(8, 16)
	password := producePassword(n)
	fmt.Println(password)
}
