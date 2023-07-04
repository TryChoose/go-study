package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"strings"
//)
//
//type JSON struct {
//	Access_token string `json:"access_token"`
//}
//
//type MESSAGES struct {
//	Touser  string `json:"touser"`
//	Toparty string `json:"toparty"`
//	Msgtype string `json:"msgtype"`
//	Agentid int    `json:"agentid"`
//	Text    struct {
//		//Subject string `json:"subject"`
//		Content string `json:"content"`
//	} `json:"text"`
//	Safe int `json:"safe"`
//}
//
//func C_Get_AccessToken(corpid, corpsecret string) string {
//	gettoken_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + corpsecret
//	//print(gettoken_url)
//	client := &http.Client{}
//	req, _ := client.Get(gettoken_url)
//	defer req.Body.Close()
//	body, _ := ioutil.ReadAll(req.Body)
//	//fmt.Printf("
//	//%q",string(body))
//	var json_str JSON
//	json.Unmarshal([]byte(body), &json_str)
//	//fmt.Printf("
//	//%q",json_str.Access_token)
//	return json_str.Access_token
//}
//
//func Send_Message(access_token, msg string) {
//	send_url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + access_token
//	//print(send_url)
//	client := &http.Client{}
//	req, _ := http.NewRequest("POST", send_url, bytes.NewBuffer([]byte(msg)))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("charset", "UTF-8")
//	resp, err := client.Do(req)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	fmt.Println("response Status:", resp.Status)
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println("response Body:", string(body))
//}
//
//func messages(touser string, toparty string, agentid int, content string) string {
//	msg := MESSAGES{
//		Touser:  touser,
//		Toparty: toparty,
//		Msgtype: "text",
//		Agentid: agentid,
//		Safe:    0,
//		Text: struct {
//			//Subject string `json:"subject"`
//			Content string `json:"content"`
//		}{Content: content},
//	}
//	sed_msg, _ := json.Marshal(msg)
//	//fmt.Printf("%s",string(sed_msg))
//	return string(sed_msg)
//}
//
//func main() {
//	touser := "C"                                               //企业号中的用户帐号，在zabbix用户Media中配置，如果配置不正常，将按部门发送。
//	toparty := "2"                                              //企业号中的部门id。
//	agentid := 1000002                                          //企业号中的应用id。
//	corpid := "ww1af03a9278c9dec4"                              //企业号的标识
//	corpsecret := "uRGUCnZ_hNBB1IaDidA88jjuW-5ya7Qd6frkXVC5_lU" ///企业号中的应用的Secret
//	accessToken := C_Get_AccessToken(corpid, corpsecret)
//	fmt.Println(accessToken)
//	content := "你好"
//	//  fmt.Println(content)
//	// 序列化成json之后，
//	//会被转义也就是变成了\n，使用str替换，替换掉转义
//	s1 := messages(touser, toparty, agentid, content)
//	msg := strings.Replace(s1, "\\", "\\", -1)
//
//	//  fmt.Println(strings.Replace(msg,"\\","\\",-1))
//	Send_Message(accessToken, msg)
//
//}
