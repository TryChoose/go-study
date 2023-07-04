package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// 设置认证信息
	auth := smtp.PlainAuth("", "yncn90@gmail.com", "Caoshuai123!", "smtp.gmail.com")

	// 设置邮件内容
	to := []string{"86127911@qq.com"}
	msg := []byte("To: 86127911@qq.com\r\n" +
		"Subject: test email\r\n" +
		"\r\n" +
		"This is a test email from Go!")

	// 发送邮件
	err := smtp.SendMail("smtp.gmail.com:465", auth, "yncn90@gmail.com", to, msg)
	if err != nil {
		fmt.Println("Error sending email: ", err)
		return
	}
	fmt.Println("Email sent successfully!")
}
