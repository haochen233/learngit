// chapter10
package main

import (
	. "fmt"
	"math/rand"
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
	ch := make(chan int)
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
