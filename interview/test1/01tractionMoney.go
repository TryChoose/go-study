package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"time"
)

type Order struct {
	gorm.Model
	Uid    int     `grom:"type:int not null" form:"Uid"`
	Weight float64 `grom:"type:double not null" form:"Weight"`
}

// 计算费用
func tranSportCosts(w float64) (m float64) {
	if w > 100.0 {
		log.Print("快递超重！")
		return
	}
	if w < 1.0 && w > 0.0 {
		w = 1
		m = 18.0
	} else {
		m = (18.0 + (w-1.0)*5.0) + (18 * (w - 1.0) * 0.01)
	}
	return
}

var DB *gorm.DB

func f0() *gorm.DB {
	dsn := "root:mysql123456@tcp(127.0.0.1:3306)/tranSportCosts?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql", dsn)
	DB = db
	return DB
}
func f1(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n) + 1
}

// 随机生成100以内的奇偶数 按照1:1进行权重分配
func f2(n int) (ans []int) {
	i, j := 0, 0
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(100) + 1
		if x%2 == 1 {
			i++
			ans = append(ans, x)
		}
		if i == n/2 {
			break
		}
	}
	for {
		x := rand.Intn(100) + 1
		if x%2 == 0 {
			j++
			ans = append(ans, x)
		}
		if j == n/2 {
			break
		}
	}
	return
}
func f4(n int) int {
	rand.Seed(time.Now().UnixNano())
	ans := f2(n)
	return ans[rand.Intn(len(ans))]
}
func f3(n int) {
	db := f0()
	fmt.Println("数据库连接成功！")
	db.AutoMigrate(&Order{})
	for i := 1; i <= n; i++ {
		order := Order{Uid: f1(i), Weight: float64(f4(n))}
		//fmt.Println(f1())
		db.Create(&order)
	}
	fmt.Println("十万条数据生成完毕！")
}
func f5() (os []Order) {
	var uid int
	db := f0()
	fmt.Printf("请输入用户id:")
	fmt.Scan(&uid)
	//db.Raw("select weight from orders where id =?", uid).Scan(&os)
	db.Where("uid=?", uid).Find(&os)
	//fmt.Println(os)
	//fmt.Println(os)
	return
}
func f6() {
	os := f5()
	var total float64
	for i := range os {
		total += tranSportCosts(os[i].Weight)
	}
	fmt.Println(total)
}
func main() {
	//f3(100000)
	f6()
}
