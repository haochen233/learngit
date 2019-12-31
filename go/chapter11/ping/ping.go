// chapter11
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
)

const (
	ICMP_RCHO_REQUEST = 8
	ICMP_RCHO_REPLY   = 0
)

//定义ICMP消息类型
type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func main() {
	if len(os.Args) != 2 {
		file := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage :%s [address]\n", file)
		os.Exit(-1)
	}

	name := os.Args[1]

	//得到IPAddr类型的地址
	ipaddr, err1 := net.ResolveIPAddr("ip", name)
	ownadr, _ := net.ResolveIPAddr("ip", "0:::")
	checkout(err1)

	//fmt.Println(ipaddr)

	//得到ip连接对象
	ipconn, err2 := net.DialIP("ip:icmp", ownadr, ipaddr)
	checkout(err2)

	//打印信息1
	if ipaddr.String() == name {
		fmt.Printf("正在 ping %s 具有28字节的数据:\n", name)
	} else {
		fmt.Printf("正在 ping %s [%s] 具有28字节的数据:\n", name, ipaddr.String())
	}

	for i := 0; i < 4; i++ {
		//new 一个新的icmp回显请求消息结构
		reqicmp := newicmp(uint16(i + 1))

		//发送开始时间
		starttime := time.Now().UnixNano()
		//发送
		_, err1 := ipconn.Write(reqicmp)
		checkout(err1)

		//设置一个buf来接收
		buf := make([]byte, 64)

		//设置超时机制
		timer := time.NewTimer(6 * time.Second)

		go func() {
			for {
				select {
				case <-timer.C:
					fmt.Println("请求超时")
					os.Exit(-1)
				}
			}
		}()
		//接收
		n, err2 := ipconn.Read(buf)
		checkout(err2)
		endtime := time.Now().UnixNano()
		costtime := float64(endtime-starttime) / float64(1e6)
		if costtime < 1 {
			fmt.Printf("来自 %s 的回复:字节=%d 时间<1ms seq=%d\n", ipaddr, n,
				i+1)
		} else {
			fmt.Printf("来自 %s 的回复:字节=%d 时间=%.2fms seq=%d\n", ipaddr, n,
				costtime, i+1)
		}
		time.Sleep(time.Second)
	}
}

//生成一个icmp报头
func newicmp(seq uint16) []byte {
	icmp := ICMP{
		ICMP_RCHO_REQUEST, //类型
		0,                 //代码
		0,                 //校验和
		0,                 //标识符
		seq,               //序列码
	}

	var buffer bytes.Buffer
	defer buffer.Reset() //清空buffer

	binary.Write(&buffer, binary.BigEndian, icmp) //把icmp头按大端序，写入buffer
	//计算出校验和
	icmp.CheckSum = checksum(buffer.Bytes())

	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)
	return buffer.Bytes()
}

//计算校验和，此时校验码为0
func checksum(data []byte) uint16 {
	length := len(data)
	sum := uint32(0)

	//将数据按16 bit累加求和，超出16位的部分会存储在sum的前16位中
	var i int
	for i = 0; i < length; i += 2 {
		//将两个8位的数据组成一个有效位为16位的uint32类型数据
		sum += uint32(data[i])<<8 + uint32(data[i+1])
	}

	if i != length {
		sum += uint32(data[length-1])
	}

	//负责将高于16位的bit位加到低16位，直到高16位为0
	for sum>>16 != 0 {
		sum = uint32(sum>>16) + uint32(sum&0xffff)
	}
	return uint16(^sum)
}

//检查错误
func checkout(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
}
