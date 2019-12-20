package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

//循环UDP服务器端
func main() {
	service := "127.0.0.1:9190"

	//将字符串类型的地址与端口转换成UDPAddr类型的
	udpaddr, err1 := net.ResolveUDPAddr("udp", service)
	CheckError(err1)

	//成功后进行监听
	for {
		udpconn, err2 := net.ListenUDP("udp", udpaddr)
		if err2 != nil {
			continue
		}
		//没有错误，进行客户端通信
		HandleClient(udpconn)
		udpconn.Close()
	}

}

//处理客户端
func HandleClient(conn *net.UDPConn) {
	//缓冲区
	var buf [1024]byte

	//先从客户端接收，再向客户端发送
	for {
		n, clntaddr, err1 := conn.ReadFromUDP(buf[0:])

		if err1 != nil {
			return
		}

		fmt.Println("Receive from client ", clntaddr.String(), string(buf[0:n]))

		//设置休眠用来测试
		time.Sleep(5 * time.Second)

		_, err2 := conn.WriteToUDP([]byte("Hello Client"), clntaddr)
		if err2 != nil {
			return
		}
	}
}

//错误检查
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
}
