package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := "192.168.174.1:9190"

	//将字符串类型的地址与端口转换成TCPAddr类型的
	tcpaddr, err1 := net.ResolveTCPAddr("tcp", service)
	CheckError(err1)

	//生成TCPListener对象
	tcplistener, err2 := net.ListenTCP("tcp", tcpaddr)
	CheckError(err2)

	//成功后进行监听
	for {
		conn, err3 := tcplistener.Accept()
		if err3 != nil {
			continue
		}

		//没有错误，进行客户端通信，每一个客户端都作为一个协程独立运行
		go HandleClient(conn)
	}

}

//处理客户端
func HandleClient(conn net.Conn) {
	//逆序调用关闭
	defer conn.Close()
	//缓冲区
	var buf [1024]byte

	//先从客户端接收，再向客户端发送
	for {
		n, err1 := conn.Read(buf[0:])

		if err1 != nil {
			break
		}

		clntaddr := conn.RemoteAddr()
		fmt.Println("Receive from client ", clntaddr.String(), string(buf[0:n]))

		//设置休眠用来测试
		time.Sleep(5 * time.Second)

		_, err2 := conn.Write([]byte("Hello Client"))
		if err2 != nil {
			break
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
