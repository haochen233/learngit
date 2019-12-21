package main

import (
	"fmt"
	"net"
	"os"
	//"reflect"
)

func main() {
	//服务器地址的字符串两种方法获取
	name, err1 := os.Hostname()
	Checkout(err1)
	//或
	//server := "127.0.0.1"

	//得到*IPAddr类型地址
	ipaddr, err2 := net.ResolveIPAddr("ip4", name)
	Checkout(err2)

	//创建IPConn对象进行连接
	ipconn, err3 := net.ListenIP("ip4:ip", ipaddr)
	Checkout(err3)

	//设置循环处理客户端
	for {
		HandleClient(ipconn)
	}
}

func Checkout(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: ", err.Error())
		os.Exit(1)
	}

}

func HandleClient(ipconn *net.IPConn) {
	defer ipconn.Close()
	var buf [1024]byte

	//接收
	n, recvipaddr, err := ipconn.ReadFromIP(buf[0:])
	if err != nil {
		return
	}

	//打印接收信息
	fmt.Println("recv from client ", recvipaddr.String()+
		string(buf[0:n]))

	//发送
	_, err2 := ipconn.WriteToIP([]byte("Hello Client!!"), recvipaddr)
	if err2 != nil {
		return
	}
}
