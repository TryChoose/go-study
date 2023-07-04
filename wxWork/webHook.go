package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type WechatRobot struct {
	Key string
}

type Message struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (robot WechatRobot) SendMessage(msg string) error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + robot.Key

	message := Message{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msg,
		},
	}

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonMsg))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	robot := WechatRobot{
		Key: "b14ccf1d-8cbf-41ac-9e1b-6701f1983ced",
	}
	err := robot.SendMessage("Hello, world!")
	if err != nil {
		panic(err)
	}
}
