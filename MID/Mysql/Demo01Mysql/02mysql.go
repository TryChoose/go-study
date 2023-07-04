package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//结构体相当于数据库里面表
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

//检查错误
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//插入
func insert(name string, age int) {
	sqlStr := `insert into users(name,age)values(?,?)`
	exec, err := db.Exec(sqlStr, name, age)
	checkErr(err)
	//_, err = exec.RowsAffected() //影响行数
	n, err := exec.LastInsertId() //插入到的哪一行
	checkErr(err)
	fmt.Println(n)
}

//删除数据
func del() {
	sqlStr := `delete from users where id=?  or id=? or id=? or id=? or id=? or id=?`
	exec, err := db.Exec(sqlStr, 9, 10, 11, 12, 13, 14)
	checkErr(err)
	n, err := exec.RowsAffected()
	checkErr(err)
	fmt.Println(n)
}

//多行读取
func queryMore(num int) {
	sqlStr := `select id ,name ,age from users where id>1 And id<?;`
	rows, err := db.Query(sqlStr, num)
	checkErr(err)
	//非常重要,必须要关闭
	defer rows.Close()
	//循环读取结果的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		checkErr(err)
		fmt.Printf("%#v\n", u)
	}
}

//单行读取
func queryOne(id int) {
	sqlStr := `select id,name,age from users where id=?;`
	var u user
	//单行查询
	/*
		 err := db.QueryRow(sqlstr, 1).Scan(&u.id, &u.name, &u.age)
				相当于 1对应的where的条件,然后将这些必需要用一个结构体接收
			必须要调用Scan方法
	*/
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	checkErr(err)
	fmt.Printf("%#v\n", u)
}

//预处理sql
func prepareInsert(name string, age int) {
	sqlStr := `insert into users(name,age)values(?,?);`
	stmt, err := db.Prepare(sqlStr)
	checkErr(err)
	defer stmt.Close()
	exec, err := stmt.Exec(name, age)
	checkErr(err)
	n, err := exec.RowsAffected()
	checkErr(err)
	fmt.Println(n)

}
func main() {
	var name string
	var age int
	err := initDB()
	checkErr(err)
	fmt.Println("数据库连接成功!")
	//queryOne(1)
	//queryMore(3)
	//for i := 0; i < 3; i++ {
	//fmt.Scan(&name)
	//fmt.Scan(&age)
	//	insert(name, age)
	//}
	//del()
	for i := 0; i < 2; i++ {
		fmt.Scan(&name)
		fmt.Scan(&age)
		prepareInsert(name, age)
	}
}
