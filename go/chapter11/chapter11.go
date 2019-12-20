// chapter11
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	/*IP地址合法性解析*/
	/*
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
			os.Exit(1)
		}

		addr := os.Args[1]
		ip := net.ParseIP(addr)

		if ip == nil {
			fmt.Println("Invaild address")
		} else {
			fmt.Println("The address is", ip.String())
		}

		os.Exit(0)
	*/

	/*域名解析*/
	/*
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
			os.Exit(1)
		}

		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name)

		if err != nil {
			fmt.Println("Invaild name")
		} else {
			fmt.Println("The address is", addr.IP.String())
		}

		os.Exit(0)
	*/

	/*获取主机地址列表信息*/

	/*	if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
			os.Exit(1)
		}

		hostname := os.Args[1]
		addrs, err := net.LookupHost(hostname)

		if err != nil {
			fmt.Println("Invaild hostname")
			os.Exit(1)
		}

		for _, v := range addrs {
			fmt.Println(v)
		}
		os.Exit(0)
	*/

	/*主机正式名查询*/

	/*	if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
			os.Exit(1)
		}

		hostname := os.Args[1]
		canonical, err := net.LookupCNAME(hostname)

		if err != nil {
			fmt.Println("Invaild hostname")
			os.Exit(1)
		}

		fmt.Println(canonical)

		os.Exit(0)
	*/

	/*查询服务器(不成功)*/
	/*	if len(os.Args) != 4 {
			fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
			os.Exit(1)
		}

		service := os.Args[1]
		protocol := os.Args[2]
		hostname := os.Args[3]
		canonical, addrs, err := net.LookupSRV(service, protocol, hostname)

		if err != nil {
			fmt.Println("Invaild hostname")
			os.Exit(1)
		}

		fmt.Println(canonical)
		fmt.Println(addrs)
		os.Exit(0)

	*/

	/*服务端口查询*/

	/*
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
			os.Exit(1)
		}

		networkType := os.Args[1]
		service := os.Args[2]
		port, err := net.LookupPort(networkType, service)

		if err != nil {
			fmt.Println("Invaild hostname")
			os.Exit(1)
		}

		fmt.Println(port)

		os.Exit(0)
	*/

	/*Dial函数举例*/
	/*
		//建立一个TCP连接，主机地址为192.168.0.1，port为5000
		conn, err := net.Dial("tcp", "192.168.0.1:5000")

		//建立一个UDP连接，主机地址为192.168.0.2，port为5001
		conn, err := net.Dial("udp", "192.168.0.2:5001")

		//建立一个ICMP连接，主机域名为www.baidu.com,省略端口号
		conn, err := net.Dial("ip4:icmp", "www.baidu.com")

		//建立一个ICMP连接，主机地址为119.75.218.77,省略端口号
		conn, err := net.Dial("ip4:1", "119.75.218.77")
	*/

	/*TCP连接地址*/
	/*
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage %s networktype addr\n", os.Args[0])
		}

		networkType := os.Args[1]
		addr := os.Args[2]

		tcpaddr, err := net.ResolveTCPAddr(networkType, addr)

		if err != nil {
			fmt.Println("ResolveTCPAddr error:", err.Error())
			os.Exit(-1)
		}

		fmt.Println("The Networktype is:", tcpaddr.Network())
		fmt.Println("The IP address is:", tcpaddr.String())

		os.Exit(0)
	*/

}
