package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// client 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("客户端发送消息：")
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("re failed,err", err)
			return
		}
		fmt.Print("收到服务器的消息：")
		fmt.Println(string(buf[:n]))
	}
}
