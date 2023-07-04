package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //驱动不能丢
	"github.com/jmoiron/sqlx"
)

//sqlx应用
type user struct {
	Id   int
	Name string
	Age  int
}

//数据库操作
var db *sqlx.DB //全局DB 很重要
//初始化数据库连接函数
func initDB() (err error) {
	dbc := "root:123456@tcp(127.0.0.1:3306)/studygo?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dbc)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(10) //设置最大连接数
	db.SetMaxIdleConns(5)  //设置闲置最大的连接数
	return
}

/*
单行查询
err := db.Get(&u, sqlstr, 1)
*/
func queryRowDemo() {
	var u user
	sqlstr := `select id,name,age from users where id=?`
	err := db.Get(&u, sqlstr, 1)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	fmt.Printf("users:%v\n", u)
}

/*
多行使用:err := db.Select(&u, sqlstr)
*/
func queryMoreDemo() {
	var u []user
	sqlstr := `select id,name,age from users`
	err := db.Select(&u, sqlstr)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	fmt.Printf("users:%v\n", u)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}
	fmt.Println("数据库连接成功！")
	queryMoreDemo()
	queryRowDemo()
}
