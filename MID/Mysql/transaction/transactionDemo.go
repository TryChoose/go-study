package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
	age  int
}

//数据库操作
var db *sql.DB //全局DB 很重要
//初始化数据库连接函数
func initDB() (err error) {
	dbc := "root:123456@(127.0.0.1:3306)/studygo"
	db, err = sql.Open("mysql", dbc)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) //设置最大连接数
	db.SetMaxIdleConns(5)  //设置闲置最大的连接数
	return
}

//开启事务
func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("事务开启失败,err:%v", err)
		return
	}
	//事务开启
	sqlStr1 := `update users set age=age-234 where id=16`
	sqlStr2 := `update users set age=age+234 where id=17`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		fmt.Println("执行sqlStr1失败,需要回滚！")
		tx.Rollback()
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Println("执行sqlStr2失败,需要回滚！")
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("事务提交失败,需要回滚！")
		tx.Rollback()
		return
	}
	fmt.Println("事务成功！")
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect mysql failed,err:%v", err)
		return
	}
	fmt.Println("数据库连接成功！")
	transactionDemo()
}
