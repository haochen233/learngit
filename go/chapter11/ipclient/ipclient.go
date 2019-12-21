package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server := "192.168.1.33"

	//将地址转换为IPAddr类型的
	ipaddr1, err1 := net.ResolveIPAddr("ip4", server)
	Checkout(err1)

	//	ipaddr2, err5 := net.ResolveIPAddr("ip4", "127.0.0.1")
	//	Checkout(err5)

	//生成连接对象进行连接
	ipconn, err2 := net.DialIP("ip4:4", nil, ipaddr1)
	Checkout(err2)
	defer ipconn.Close()

	//连接成功后进行通信
	var buf [1024]byte

	//发送
	_, err3 := ipconn.Write([]byte("Hello Server!"))
	Checkout(err3)

	//接收
	n, _, err4 := ipconn.ReadFromIP(buf[0:])
	Checkout(err4)

	fmt.Println("Recv server ", ipaddr1.String(), string(buf[0:n]))
}

func Checkout(err error) {
	if err != nil {
		fmt.Println("fatal error: ", err.Error())
		os.Exit(-1)
	}

}
