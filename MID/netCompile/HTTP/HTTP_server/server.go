package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTP server 服务端开发

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./01.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}
func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求，参数都放在URL上（query param）,请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL中的参数
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	//fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) //服务器打印客户端发来的请求的body
	w.Write([]byte("ok"))
}
func main() {
	//回调函数
	http.HandleFunc("/index/", f1)
	http.HandleFunc("/in/", f2)
	http.ListenAndServe("127.0.0.1:20000", nil)

}
