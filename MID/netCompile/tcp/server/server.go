package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	//3.与客户端通信
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		fmt.Print("请回复客户端：")
		sendMsg := bufio.NewReader(os.Stdin)
		input, _ := sendMsg.ReadString('\n')
		msg := strings.Trim(input, "\r\n")
		conn.Write([]byte(msg)) //发送数据
	}

}

//tcp server端
func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	//2.等待别人来跟我连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		go process(conn)
	}

}
