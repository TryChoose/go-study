package main

//数组
func main() {
	////先声明后初始化
	////数组的声明
	//var a1 [3]int
	////初始化数组
	//a1=[3]int{1,2,3}
	//fmt.Println(a1)
	////同时声明并初始化
	//var a2=[3]int{1,2,4}
	//fmt.Println(a2)
	//var s1=[3]string{"武汉","上海","北京"}
	//fmt.Println(s1)
	////自行推断初始值和数组长度
	//var a3=[...]int{2,3,4,5,6}
	//fmt.Println(a3)
	//fmt.Println(len(a3))
	//fmt.Printf("a3:%T\n",a3)
	////通过索引方式来初始化数组
	//a:=[...]int{0:1,4:5,7:1}
	//fmt.Println(a)
	//数组的遍历
	//var a=[...]int{1,2,3,4,5,6,7,8,9,0}
	//for i := 0; i <len(a) ; i++ {
	//	fmt.Printf("%d ",a[i])
	//}
	//fmt.Println()
	//for _,v:=range a{
	//	fmt.Printf("v:%d\n",v)
	//}
	//二维数组
	//a:=[3][2]int{
	//	{1,2},
	//	{2,3},
	//	{3,5},
	//}
	//for i:=0;i<3;i++ {
	//	for j:=0;j<2;j++ {
	//		fmt.Printf("%d\t",a[i][j])
	//	}
	//	fmt.Println()
	//}
	//for _,v1:=range a{
	//	for _,v2:=range v1{
	//		fmt.Printf("%d\t",v2)
	//	}
	//	fmt.Println()
	//}
	//[[1 2] [6 3] [5 4]]
	//var a2  [3][2]string
	//a2=[3][2]string{
	//	{"1","2"},
	//	{"6","3"},
	//	{"5","4"},
	//}
	//fmt.Println(a2)
	//a:=[...]int{1,3,5,7,8}
	//sum:=0
	//for _,v:=range a{
	//	//fmt.Println(v)
	//	sum+=v
	//}
	//fmt.Println(sum)

	//for i:=0;i<len(a);i++{
	//	for j:=i+1;j<len(a);j++{
	//		if a[i]+a[j]==8 {
	//			fmt.Printf("(%d,%d) ",i,j)
	//		}
	//	}
	//}
	//a1:=[...][2]string{
	//	{"长沙","武汉"},
	//	{},
	//	{},
	//}
	//fmt.Println(a1)
	//fmt.Println(len(a1))

}
