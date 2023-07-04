package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
重点：随机数去重,清屏
难点: 随机数去重
*/
func main() {
	fmt.Println("您好,现在您有十秒钟的时间记住当前物品所对应的序号")
	tings := []string{"苹果", "香蕉", "梨子", "菠萝", "柚子", "西瓜", "葡萄"}
	for i, v := range tings {
		fmt.Println(i, v)
	}
	time.Sleep(10 * time.Second)
	//清屏指令
	fmt.Print("\033[H\033[2J")
	//cmd := exec.Command("cmd", "/c", "cls")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	//去重
	f1 := func(n, len int) (res []int) {
		count := 0
		for {
			rand.Seed(time.Now().Unix())
			num := rand.Intn(len)
			if count == n {
				break
			}
			exist := false
			for _, v := range res {
				if v == num {
					exist = true
					break
				}
			}
			if !exist {
				res = append(res, num)
				count++
			}
		}
		return
	}
	rs := []string{}
	res := f1(3, len(tings))
	for _, v := range res {
		rs = append(rs, tings[v])
	}
	fmt.Println(rs)
	//输出序号
	var x, count int
	for i, v := range rs {
		fmt.Print(v, ":")
		//输入数字
		fmt.Scan(&x)
		if x == res[i] {
			count++
		}
	}
	fmt.Println("答对了：", count)
}
