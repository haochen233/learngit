package main

import (
	. "fmt"
	//"runtime"
	. "time"
)

func main() {
	/*
	   Println(LoadLocation("Asia/Shanghai"))
	   	t := Now()
	   	Println(t)
	   	Println(LoadLocation(""))
	   	t1, _ := ParseInLocation("2006-01-02 15:04:05", "2019-12-27 10:39:26", Local)

	   	Println(t1)
	*/

	//时区与当前时区不同，所以一般用ParseInLocation来解析时间
	/*	t2, _ := Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
		Println(t2)
	*/
	/*
		Println(Hour)
		Println(Minute)
		Println(Second)
		Println(Millisecond)
		Println(Nanosecond)
	*/
	//把Duration类型的单位向下取整 Method Truncate（只能舍去）
	//Println(t1.Truncate(Hour))

	//Println(t1.Truncate(Minute))
	//Println(t1.Truncate())

	//把Duration类型的单位向下或向上取整 Method Round(根据四舍五入，既能舍也能入)
	//Println(t1.Round(Hour))

	/*time.After模拟超时*/

	/*	timer := NewTimer(2 * Second)
		c := make(chan int)

		go func() {
			Sleep(1 * Second)
			if timer.Reset(6 * Second) {
				Println("Reset timer")
			}
			Sleep(5 * Second)
			<-c
		}()

		select {
		case c <- 1:
			Println("channel...")
		case <-timer.C:
			close(c)
			Println("timeout")
		}
	*/

	timer := NewTimer(2 * Second)
	ticker := NewTicker(2 * Second)
	c := make(chan int)

	go func() {
		if timer.Reset(9 * Second) {
			Println("Reset timer")
		}
		Sleep(5 * Second)
		<-c
	}()

	for {
		select {
		case c <- 1:
			Println("channel...")
		case <-timer.C:
			ticker.Stop()
			Println("用Timer关闭Ticker")
			close(c)
			Println("Timer timeout")
			return
		case <-ticker.C:
			Println("Ticker timeout")
		}
	}
}
