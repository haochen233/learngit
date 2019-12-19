// chapter10
package main

import (
	. "fmt"
	"math/rand"
	"runtime"
	"time"
	//"reflect"
)

func main() {
	/*Channel的声明和初始化*/
	/*
		var ch1 chan int
		ch2 := make(chan int)
		var chs1 []chan int
		chs2 := make([]chan int, 8)
		for i := 0; i < 10; i++ {
			go test(i)
		}

		t1 := reflect.TypeOf(ch1)
		t2 := reflect.TypeOf(ch2)
		ts1 := reflect.TypeOf(chs1)
		ts2 := reflect.TypeOf(chs2)

		Println(t1.Kind())
		Println(t2.Kind())
		Println(ts1)
		Println(ts1.Name(), ts1.Kind())
		Println(ts2)
		Println(ts2.Name(), ts2.Kind())
	*/

	/*数据的发送与接收*/
	/*	chs := make([]chan int, 10)
		for i, _ := range chs {
			chs[i] = make(chan int)
			go Test(chs[i])
		}

		for _, ch := range chs {
			value := <-ch
			Println(value)
		}
	*/

	/*Channel的迭代操作*/

	/*	ch := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				ch <- i
			}
			//记住关闭，否则下面range遍历出问题
			close(ch)
		}()


		for value := range ch {
			Print(value, ",")
		}

		//返回false表示已关闭
		_, ok := <-ch
		if !ok {

			Println("close successed")
		}
	*/

	/*
		ch := make(chan int)
		lock := make(chan bool)

		go Send(ch)
		go Recv(ch, lock)
		//两个

		ok := <-lock
		if ok {
			Println("channel normal closed")
		}
	*/

	//创建Buffer为2的int 型Channel
	/*	ch := make(chan int, 2)
		lock := make(chan bool)

		for i := 0; i < 10; i++ {
			go Worker(ch, lock, i)
		}

		<-lock
	*/

	/*select机制使用举例*/
	/*
		ch1 := make(chan int)
		ch2 := make(chan string)
		lock := make(chan bool)

		go func() {
			for {
				select {
				case v, ok := <-ch1:
					if !ok {
						lock <- true
						//break
					}
					Println("ch1:", v)

				case v, ok := <-ch2:
					if !ok {
						lock <- true
						//break
					}
					Println("ch2:", v)
				}
			}
		}()

		ch1 <- 66
		ch2 <- "中国"
		ch1 <- 88
		ch2 <- "niubi"
		ch2 <- "2020"
		ch2 <- "全面小康"
		close(ch1)
		close(ch2)
		<-lock
	*/

	/*利用select实现超时机制*/
	/*	ch := make(chan int)
		timeout := make(chan bool)

		go func() {
			time.Sleep(5 * time.Second)
			timeout <- true
		}()

		select {
		case <-ch:
		case <-timeout:
			Println("timeout...")
		}
	*/

	/*After函数实现超时机制*/
	/*	ch := make(chan int)
		lock := make(chan bool)

		go func() {
			for {
				select {
				case <-ch:
				case <-time.After(5 * time.Second):
					Println("timeout")
					lock <- true
				}
			}
		}()

		<-lock
	*/

	/*Gosched出让时间片*/
	/*
		//Goroutine1
		go func() {
			for i := 0; i < 100; i++ {
				if i == 2 {
					runtime.Gosched()
				}

				Println("Goroutine 1:", i)
			}
		}()

		//Goroutine2
		go func() {
			Println("Goroutine2")
		}()

		time.Sleep(5 * time.Second)
	*/

	/*统计CPU核心数和任务数*/
	/*	Println("CPU number:", runtime.NumCPU())
		Println("Goroutines start:", runtime.NumGoroutine())

		ch := make(chan int)

		//创建一些Goroutine
		for i := 0; i < 5; i++ {
			go func(n int) {
				<-ch
				Println(n, runtime.NumGoroutine())
				ch <- 1
			}(i)
		}

		ch <- 1

		time.Sleep(5 * time.Second)
		Println("Goroutines over:", runtime.NumGoroutine())
	*/

	/*Goexit 强行终止Goroutine*/
	/*
		go func() {
			defer Println("Goroutine1 defer...")
			for i := 0; i < 10; i++ {
				if i == 5 {
					runtime.Goexit()
				}

				Println("Goroutine1:", i)
			}
		}()

		go func() {
			Println("Goroutine2")
		}()

		time.Sleep(5 * time.Second)
	*/
}

func Worker(sem chan int, lock chan bool, id int) {
	//填入缓冲区
	sem <- 1
	Println(time.Now().Unix(), "Goid:", id)
	time.Sleep(1 * time.Second)
	<-sem
	if id == 9 {
		lock <- true
	}

}

func Recv(ch <-chan int, lock chan<- bool) {
	for value := range ch {
		Println(value)
	}

	lock <- true
}

func Send(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
}

func Test(ch chan int) {
	ch <- rand.Int() //向channel中写入一个随机数
	Println("Go...")
}

/*
func test(n int) {
	Println("Goroutine", n)

}
*/
