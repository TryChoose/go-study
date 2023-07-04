package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库实例
/**w
 1.首先连接数据库
  mysql -uroot -p
	123456
2.show databases;
3.create database studygo;
*/
func main() {
	//数据库信息
	dbc := "root:123456@tcp(127.0.0.1:3306)/studygo"
	//用户名:密码@连接的端口/数据库名称
	//连接数据库
	db, err := sql.Open("mysql", dbc) //验证数据库信息格式是否正确
	if err != nil {
		fmt.Printf("dbc %s invalid,err:%v\n", dbc, err)
		return
	}
	err = db.Ping() //尝试进行连接数据库
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dbc, err)
		return
	}
	fmt.Println("连接数据库成功!")
}
