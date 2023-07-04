package main

import (
	"fmt"
	"net"
)

//UDP服务端
func main() {
	//1.连接的步骤
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000})
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	defer udpConn.Close()
	//2.读取数据，写入数据
	for {
		//接收数据
		var data [1024]byte
		n, addr, err := udpConn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("readFromUDP failed,err:", err)
			return
		}
		fmt.Println(string(data[:n]))
		//发送数据
		_, err = udpConn.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed,err", err)
			return
		}
	}
}
