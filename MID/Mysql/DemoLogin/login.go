package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

/*
 登录案例
需求：
实现登录功能和注册功能,修改密码

1.使用数据库存放数据 用户数据

*/

type account struct {
	id       int
	user     string
	password string
}

var (
	db  *sql.DB
	a   account
	zh  string
	pwd string
)

//登录函数
/*
	主要逻辑:
	首先需要判断输入的账号和密码,是否在数据库有对应的,如果有对应的则登陆成功
	相当于查询操作
*/
func login(zh string, pwd string) (flag bool) {
	flag = true
	sqlStr := `select id,name,pwd from acc where name=? and pwd=?;`
	err := db.QueryRow(sqlStr, zh, pwd).Scan(&a.id, &a.user, &a.password)
	if err != nil {
		fmt.Printf(" login failed ,err:%v\n", err)
		flag = false
	}
	return
}

//注册函数
/*
	1.需要判断一下输入的数据是否在数据库中存在
	2.相当于往数据库表添加一条数据
*/
func register(zh, mm string) (flag bool) {
	//1.判断一下是否存在
	sqlSearch := `select name from acc where name=? and flag=0`
	err := db.QueryRow(sqlSearch, zh).Scan(&a.user)
	if err != nil {
		//fmt.Printf("注册账号不存在,err:%v", err)
		flag = true
		//err不为空 说明查询不到数据 所以可以往里插入数据
		sqlStr := `insert into acc (name,pwd)values(?,?);`
		_, err = db.Exec(sqlStr, zh, mm)
		if err != nil {
			fmt.Printf("register account failed,err:%#v\n", err)
			flag = false
		}
	} else {
		fmt.Print("账户信息已经存在,")
	}
	return
}

/*
注销函数
首先判断是否存在账户和密码在数据表中,之后进行删除该数据 更新对应的默认值 实现逻辑删除
*/
func logOff(zh, mm string) (flag bool) {
	sqlSearch := `select name from acc where name=? and flag=0`
	err := db.QueryRow(sqlSearch, zh).Scan(&a.user)
	if err != nil {
		fmt.Println("该用户不存在 无法注销")
		return
	} else {
		sqlUpdate := `update acc set flag=1 where name=? and pwd=?`
		_, err := db.Exec(sqlUpdate, zh, mm)
		if err != nil {
			fmt.Printf("删除失败,err:%v\n", err)
			return
		} else {
			return true
		}
	}
}

/*
修改密码函数
首选需要判断在表中，是否存在该账户(存在且标识位为0)
之后进行更改
*/
func alterMM(zh, mm string) (flag bool) {
	sqlSearch := `select name from acc where name=? and flag=0`
	err := db.QueryRow(sqlSearch, zh).Scan(&a.user)
	if err != nil {
		fmt.Println("该用户不存在 无法修改密码！")
		return
	} else {
		var newmm string
		fmt.Printf("请输入新的密码：")
		fmt.Scan(&newmm)
		sqlUpdate := `update acc set pwd=? where name=? and flag=0`
		_, err := db.Exec(sqlUpdate, newmm, zh)
		if err != nil {
			fmt.Printf("修改失败,err%v\n", err)
			return
		} else {
			flag = true
			return
		}
	}
}

//检查错误函数
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}

//数据库登录初始化函数
func initDB() (err error) {
	dbc := "root:123456@tcp(127.0.0.1:3306)/login"
	db, err = sql.Open("mysql", dbc)
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	return
}

//输入函数
func f1() {
	fmt.Print("请输入账号：")
	fmt.Scan(&zh)
	fmt.Print("请输入密码：")
	fmt.Scan(&pwd)

}
func main() {
	err := initDB()
	checkErr(err)
	//fmt.Println("数据库连接成功")
	fmt.Println("===============菜单===============")
	fmt.Printf("\t   ======1.登录账号======\n" +
		"\t   ======2.注册账号======\n" +
		"\t   ======3.注销账号======\n" +
		"\t   ======4.修改密码======\n" +
		"\t   ======5.退出程序======\n")
	fmt.Print("请输入执行的操作序号：")
	var op int
	fmt.Scan(&op)
	switch op {
	case 1:
		f1()
		flag := login(zh, pwd)
		if flag {
			fmt.Println("登陆成功!")
		} else {
			fmt.Println("登录失败！")
		}
	case 2:
		f1()
		flag := register(zh, pwd)
		if flag {
			fmt.Println("注册成功!")
		} else {
			fmt.Println("注册失败！")
		}
	case 3:
		f1()
		flag := logOff(zh, pwd)
		if flag {
			fmt.Println("删除成功！")
		} else {
			fmt.Println("删除失败！")
		}
	case 4:
		f1()
		flag := alterMM(zh, pwd)
		if flag {
			fmt.Println("修改密码成功！")
		} else {
			fmt.Println("修改密码失败！")
		}
	case 5:
		os.Exit(1)
	}
}
