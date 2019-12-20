package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s addr\n", os.Args[0])
		os.Exit(-1)
	}

	addr := os.Args[1]

	//进行地址转换
	udpaddr, err1 := net.ResolveUDPAddr("udp", addr)
	CheckError(err1)

	//建立连接
	udpconn, err2 := net.DialUDP("udp", nil, udpaddr)
	CheckError(err2)
	defer udpconn.Close()

	//进行通信
	var buf [1024]byte

	//发送
	_, err3 := udpconn.Write([]byte("Hello UDPserver"))
	CheckError(err3)

	//接收
	n, servaddr, err4 := udpconn.ReadFromUDP(buf[0:])
	CheckError(err4)

	fmt.Println("receive from server ", servaddr.String(), string(buf[0:n]))
}

//错误检查
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
}
