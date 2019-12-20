package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s servaddr", os.Args[0])
		os.Exit(-1)
	}

	servaddr := os.Args[1]

	//将字符串格式的服务器地址信息转换成TCPAddr类型
	tcpaddr, err1 := net.ResolveTCPAddr("tcp", servaddr)
	CheckError(err1)

	//建立TCP连接
	conn, err2 := net.DialTCP("tcp", nil, tcpaddr)
	CheckError(err2)

	//接下来与服务器进行通信
	var buf [1024]byte
	saddr := conn.RemoteAddr()

	//发
	_, err4 := conn.Write([]byte("Hello Server!"))
	CheckError(err4)

	//收
	n, err5 := conn.Read(buf[0:])
	CheckError(err5)

	fmt.Println("Receive form server ", saddr.String(), string(buf[0:n]))

	//关闭连接
	conn.Close()
	os.Exit(1)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
}
