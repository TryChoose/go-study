package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//复用client连接，适用于请求比较频繁
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	//resp, err := http.Get("http://127.0.0.1:20000/in/?name=ABC&age=20")
	//if err!=nil{
	//	fmt.Println("get url failed,err:",err)
	//	return
	//}
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:20000/in/")
	data.Set("name", "李四")
	data.Set("age", "20")
	encode := data.Encode()
	fmt.Println(encode)
	urlObj.RawQuery = encode
	request, err := http.NewRequest("GET", urlObj.String(), nil)
	//resp, err := http.DefaultClient.Do(request)
	//if err!=nil{
	//	fmt.Println("get url failed,err:",err)
	//	return
	//}
	//禁用 KeepAlive的Client ，请求不是特别频繁，一天两次，用完就关闭该连接
	//tr := &http.Transport{
	//	DisableKeepAlives: true,
	//}
	//client := http.Client{
	//	Transport: tr,
	//}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("get url failed,err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) //在客户端读出服务器响应的body
	if err != nil {
		fmt.Println("read resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(b))
}
