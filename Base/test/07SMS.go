package main

import (
	"fmt"
	"os"
)

var (
	stu map[int64]*student
)

type student struct {
	id    int64
	name  string
	score float64
}

//student的构造函数
func newStudent(id int64, name string, score float64) *student {
	return &student{id: id, name: name, score: score}
}

/*
学生管理系统函数版
1.查看学生
2.添加学生
3.删除学生
*/
func show() {
	for k, v := range stu {
		fmt.Printf("学生的学号是:%d,姓名是:%s,成绩是:%.2f\n", k, v.name, v.score)
	}
}
func add() {
	var (
		id    int64
		name  string
		score float64
	)
	fmt.Print("请输入学号：")
	fmt.Scan(&id)
	if _, ok := stu[id]; ok {
		fmt.Print("该生已经存在,请重新选择\n")
		//add()
		return
	}
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入成绩：")
	fmt.Scan(&score)

	stu[id] = newStudent(id, name, score)
}
func del() {
	var delNub int64
	fmt.Print("请输入要删除的学生学号:")
	fmt.Scan(&delNub)
	if _, ok := stu[delNub]; ok {
		delete(stu, delNub)
	}
}
func main() {

	stu = make(map[int64]*student, 50)
	for {
		//1.欢迎界面
		fmt.Println("欢迎进入学生管理系统！！！")
		//2.菜单
		fmt.Print(`
		1.查看学生
		2.添加学生
		3.删除学生
		4.退出系统
	`)
		//3.输入的操作的编号
		fmt.Print("请输入您要操作的编号:")
		var o int64
		fmt.Scan(&o)
		fmt.Printf("您选择的是%d号功能.\n", o)
		switch o {
		case 1:
			show()
		case 2:
			add()
		case 3:
			del()
		case 4:
			fmt.Println("已退出该系统")
			os.Exit(1)
		default:
			fmt.Println("暂未此功能,敬请期待....")
			break
		}
	}
}
