package main

import (
	"fmt"
	"os"
)

/*
代码逻辑:通过管理员调用方法，查看/添加/删除学生信息
*/
//学生管理系统方法版

type manger struct {
	id       int
	password string
}

type students struct {
	id   int
	name string
	age  int
}

var stu1 map[int]*students
var managers map[int]string

//学生的构造函数
func newStudents(id int, name string, age int) *students {
	return &students{id: id, name: name, age: age}
}

//管理员的构造函数
func newManger(id int, password string) *manger {
	return &manger{id: id, password: password}
}

//管理员登陆函数
func mangerSign() *manger {
	fmt.Println("请登录管理员账号!!")
	var id int
	var psw string
	fmt.Print("请输入账号:")
	fmt.Scan(&id)
	fmt.Print("请输入密码:")
	fmt.Scan(&psw)
	m := newManger(id, psw)
	managers[id] = psw
	//登录之后还需要记住当前账号和密码 或者游客/管理员登陆 不着急实现
	return m
}

/*
学生管理系统函数版
1.查看学生
2.添加学生
3.删除学生
*/
//游客展示学生信息函数(都可以调用)
func showS() {
	if len(stu1) == 0 {
		fmt.Println("学生系统里面没有学生信息！")
	}
	for k, v := range stu1 {
		fmt.Printf("学生的id：%d,姓名：%s,年龄:%d\n", k, v.name, v.age)
	}
}

//查看学生方法
func (m manger) showS() {
	if len(stu1) == 0 {
		fmt.Println("学生系统里面没有学生信息！")
	}
	for k, v := range stu1 {
		fmt.Printf("学生的id：%d,姓名：%s,年龄:%d\n", k, v.name, v.age)
	}
	fmt.Println("查看成功！")
}

//添加学生方法
func (m manger) addS() {
	var (
		id   int
		name string
		age  int
	)
	fmt.Println("请输入添加的个人信息")
	fmt.Print("学号：")
	fmt.Scan(&id)
	fmt.Print("姓名：")
	fmt.Scan(&name)
	fmt.Print("年龄：")
	fmt.Scan(&age)
	stu1[id] = newStudents(id, name, age)
	fmt.Println("添加成功!")
}

//修改学生方法
func (m manger) updateS() {
	//1.需要获取该学生是否存在
	var sID int
	fmt.Print("请输入学生的学号：")
	fmt.Scan(&sID)
	//2.查看学生初始信息
	value, ok := stu1[sID]
	if !ok {
		fmt.Print("不存在有该学生!\n")
	}
	fmt.Printf("学生的学号：%d,姓名：%s,年龄：%d\n", value.id, value.name, value.age)
	//3.修改学生新的信息
	var (
		newName string
		newAge  int
	)
	fmt.Print("请输入学生的新姓名：")
	fmt.Scan(&newName)
	fmt.Print("请输入学生的新年龄：")
	fmt.Scan(&newAge)
	stu1[sID].name = newName
	stu1[sID].age = newAge
	fmt.Print("修改成功！")

}

//删除学生方法
func (m manger) delS() {
	var delId int
	fmt.Print("请输入要删除学生的id:")
	fmt.Scan(&delId)
	delete(stu1, delId)
	fmt.Println("删除成功！")
}
func menu(m *manger) {
	for {
		//菜单
		fmt.Print(`
		1.查看学生
		2.添加学生
		3.修改学生
		4.删除学生
		5.退出系统
	`)
		//选择需要操作的数字
		var a int
		fmt.Scan(&a)
		switch a {
		case 1:
			//查看学生
			m.showS()
		case 2:
			//添加学生
			m.addS()
		case 3:
			m.updateS()
		case 4:
			//删除学生
			m.delS()
		case 5:
			//退出系统
			os.Exit(1)
		default:
			fmt.Println("输入的操作不存在，敬请期待...")
		}
	}
}
func main() {

	//开辟内存空间
	stu1 = make(map[int]*students, 40) //开辟了空间不等于nil
	managers = make(map[int]string, 10)
	fmt.Println("欢迎来到学生管理系统方法版!!!")
	fmt.Print("登录请输入1不登录请输入0：")
	var x int
	fmt.Scan(&x)
	if x == 1 {
		m := mangerSign()
		menu(m)
		fmt.Println("登陆成功")

	} else if x == 0 {
		showS()
	}

}
