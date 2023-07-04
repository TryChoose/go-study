package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币 //69 101
b. 名字中每包含1个'i'或'I'分2枚金币 73 105
c. 名字中每包含1个'o'或'O'分3枚金币 79 111
d: 名字中每包含1个'u'或'U'分4枚金币 84 117
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

/*分析
1.使用map存放人名和对应的金额
2.使用字符串切片就算出每个人名该分配多少金币
3.存放在map里面
4.输出map
*/
func dispatchCoin() {
	var (
		count = 0
		coins = 50
		users = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	)
	m1 := make(map[string]int, len(users))
	for _, v := range users {
		//fmt.Println(i,v)
		for _, v1 := range v {
			switch v1 {
			case 69, 101:
				count += 1
			case 73, 105:
				count += 2
			case 79, 111:
				count += 3
			case 84, 117:
				count += 4
			}
			//case v1=="e"||v1=="E":
			//	count+=1

		}
		m1[v] = count
		coins -= count
		count = 0
	}
	for s, i := range m1 {
		fmt.Println(s, i)
	}
	fmt.Println(coins)
}
func main() {
	dispatchCoin()
}
