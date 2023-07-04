package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "www.baidu.com"
	ip, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range ip {
		fmt.Printf("%v,%T\n", v, v)
	}
}
