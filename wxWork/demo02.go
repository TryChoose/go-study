package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSON struct {
	Access_token string `json:"access_token"`
}

func Get_AccessToken(corpid, corpsecret string) string {
	gettoken_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + corpsecret
	//print(gettoken_url)
	client := &http.Client{}
	req, _ := client.Get(gettoken_url)
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	//fmt.Printf("
	//%q",string(body))
	var json_str JSON
	json.Unmarshal([]byte(body), &json_str)
	//fmt.Printf("
	//%q",json_str.Access_token)
	return json_str.Access_token
}
func main() {
	// 企业微信发送消息接口
	corpid := "ww1af03a9278c9dec4"                              //企业号的标识
	corpsecret := "uRGUCnZ_hNBB1IaDidA88jjuW-5ya7Qd6frkXVC5_lU" ///企业号中的应用的Secret
	accessToken := Get_AccessToken(corpid, corpsecret)
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accessToken

	// 请求参数
	params := make(map[string]interface{})
	params["touser"] = "@all"
	params["msgtype"] = "text"
	params["agentid"] = 1000002
	params["text"] = map[string]string{"content": "测试企业微信自动化发送消息"}

	// 将参数转换为json
	paramsJson, _ := json.Marshal(params["text"])
	// 发送请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(paramsJson))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 读取响应
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(respBytes))
}
