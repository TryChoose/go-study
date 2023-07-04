package main

import . "fmt"

//map

func main() {
	////map的定义：map[KeyType]ValueType
	//var m1 map[int]int //map必须初始化
	//m1=make(map[int]int,10)
	//m1[1]=10
	//m1[2]=20
	//fmt.Println(m1)
	//m1 := make(map[int]int, 10)
	//m1[0] = 20
	//m1[1] = 30 //会覆盖原来的值
	//fmt.Println(m1)
	//fmt.Printf("%T\n", m1)
	//
	//m2 := map[int]int{
	//	1: 20,
	//	2: 30,
	//}
	//fmt.Println(m2)
	////make判断键是否存在
	//value, ok := m2[3]
	//if ok == false {
	//	fmt.Println(value) //0
	//} else {
	//	fmt.Println("无效key")
	//}

	//m1 := make(map[string]int, 10)
	//m1["张三"] = 100
	//m1["李四"] = 200
	//m1["王五"] = 300
	////遍历key和value
	//for k, v := range m1 {
	//	fmt.Println(k, v)
	//}
	////只遍历key 默认只有一个变量时 遍历的是key
	//for k := range m1 {
	//	fmt.Println(k)
	//}
	////遍历value
	//for _,v := range m1 {
	//	fmt.Println(v)
	//}
	////delete函数 删除
	//delete(m1,"王五")
	//fmt.Println(m1)
	//按照指定的顺序进行遍历:根据键排序
	//rand.Seed(time.Now().UnixNano()) //初始化随机种子
	//stu:=make(map[string]int,100)
	//for i:=0;i<10;i++ {
	//	//键重复了 然后覆盖了
	//	key:=fmt.Sprintf("%d",i)
	//	v:=rand.Intn(10)
	//	stu[key]=v
	//}
	////stu["张三3"]=100
	////stu["李三1"]=200
	////stu["王五3"]=100
	//var keys=make([]string,0,200)
	//for key := range stu {
	//	keys=append(keys,key)
	//}
	//sort.Strings(keys)
	//for _,key:=range keys{
	//	fmt.Println(key,stu[key])
	//}

	//元素（key）为map类型的切片
	var mapSlice = make([]map[string]string, 3, 10) //对整个类型进行初始化
	//mapSlice := make([]map[string]int, 3, 10)
	Printf("%T\n", mapSlice)
	mapSlice[0] = make(map[string]string, 10) //对map[0]进行初始化
	Printf("%T\n", mapSlice[0])
	mapSlice[1] = make(map[string]string, 10) //对map[1]进行初始化
	mapSlice[2] = make(map[string]string, 10) //对map[2]进行初始化
	mapSlice[0]["你好"] = "Hello"               //整个切片赋值
	mapSlice[1]["世界"] = "World"
	mapSlice[2]["!"] = "!"
	//for i := 0; i < 3; i++ {
	//	//	mapSlice[i] = make(map[string]string, 10)
	//	//}
	//word := ""
	//for i := 0; i < 3; i++ {
	//	fmt.Scanf("%s", &word)
	//	mapSlice[i][word] = i
	//}
	Println(mapSlice)
	for _, m := range mapSlice {
		Println(m)
	}
	//值为切片类型的map
	//var m1 = make(map[string][]int,10)
	//m1["北京"]=make([]int,0)
	//m1["北京"]=append(m1["北京"],10,20,30)
	////m1["北京"]=[]int{10,20,30}
	//fmt.Println(m1)

}
