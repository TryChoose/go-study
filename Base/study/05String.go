package main

import "fmt"

func main() {
	//字符串拼接
	/*
		name:="张三"
		age:="19"
		s1:=name+age
		fmt.Println(len(s1))
		fmt.Println(s1)
		s2 := fmt.Sprintf("%s%s",name,age)
		fmt.Println(s2)
		//分割
		path:="E:\\DemoProject\\study.golang.cn\\study"
		ret := strings.Split(path, "\\")
		fmt.Println(ret)
		//包含
		fmt.Println(strings.Contains(path,"cn"))
		fmt.Println(strings.Contains(path,"cm"))
		//前缀
		fmt.Println(strings.HasPrefix(path,"E"))
		//后缀
		fmt.Println(strings.HasSuffix(path,"y"))
		//字符串子串
		s4:="abcdefac"
		fmt.Println(strings.Index(s4,"d"))
		fmt.Println(strings.LastIndex(s4,"a"))
		fmt.Println(strings.Join(ret,"/"))
		s:="HelloWorldA你好"
		for i := 0; i <len(s) ; i++ {
			fmt.Printf("%v(%c)",s[i],s[i]) //ASCII 中文会乱码
		}
		fmt.Println()
		for _,r:=range s{
			fmt.Printf("%v(%c)",r,r) //rune 不会乱码
		}
		fmt.Println()
	*/
	/*
		修改字符串操作
		1.将字符串转为[]rune 或者[]byte
		2.对某为字符进行修改
		3.最后转为string
	*/
	s1 := "王者荣耀"
	fmt.Println(s1)
	r1 := []rune(s1)
	r1[0] = '无'
	fmt.Println(string(r1))
	s2 := "Hello"
	fmt.Println(s2)
	b1 := []byte(s2)
	b1[0] = 'W'
	fmt.Println(string(b1))
}
