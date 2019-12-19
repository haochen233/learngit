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

}
