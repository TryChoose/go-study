package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//UDP客户端
func main() {
	//连接服务器
	dialUDP, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000})
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer dialUDP.Close()
	var rData [1024]byte
	//从终端上面输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		input, _ := inputReader.ReadString('\n')
		//转换一下换行和空格
		inputInfo := strings.Trim(input, "\r\n")
		dialUDP.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("发送数据失败：", err)
			return
		}
		//接收数据
		n, _, err := dialUDP.ReadFromUDP(rData[:])
		if err != nil {
			fmt.Println("接收数据失败,err", err)
			return
		}
		fmt.Println("接收到数据:", string(rData[:n]))
	}
}
