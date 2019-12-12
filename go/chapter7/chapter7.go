package main

import (
	"errors"
	. "fmt"
	//"mymath"
	//"time"
)

/*无参无返回值*/
func f1() {
	Println("Hello function")
}

/*有参无返回值*/
func f2(a, b int, c string) {
	Println(a, b, c)
}

/*有参有一个返回值*/
func f3(a, b int) int {
	return a + b
}

/*有参有多个返回值*/
func f4(a, b int) (ret float32, err error) {
	if b == 0 {
		err = errors.New("overflow!")
		return
	} else {
		return float32(a) / float32(b), nil
	}
}

func main() {

	//f1()
	//f2(1, 2, "666")
	//Println("a+b=", f3(32, 34))
	//Println(f4(10, 0))

	/*调用时间函数获取当前日期*/

	//获取时间戳
	/*	t := time.Now()

		Println(t)
		Println(t.String())

		//按 年月日格式输出
		Println(t.Format("2006 年 01 月 02 日"))

		//输出星期几
		Println(t.Weekday())
		Println(t.Weekday().String())
	*/

	/*使用mymath中的函数*/
	/*
		Println(mymath.Add(1, 2))
		Println(mymath.Sub(100, 34))
		Println(mymath.Mult(3, 22))
		Println(mymath.Div(132, 2))
	*/

	/*
		var i int = 66
		var p *int = &i

		Println(i)

		f5(p)

		Println(i)
	*/

	/*	var arr1 = [5]int{1, 2, 3, 4, 5}

		var arr2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		var arr3 [5]int = [5]int{1, 2, 3, 4, 5}

		var arr4 = [5]int{1: 66, 4: 88}

		f6(arr1)

		Println(arr1[0])
		Println(arr2)
		Println(arr3)
		Println(arr4)
	*/

	/*	var arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		var slice1 []int = arr1[0:5]

		var slice2 = arr1[5:]

		Println(arr1[0])
		f7(slice1)
		Println(arr1[0])

		Println(len(slice2))

		Println(slice2[0])
		f7(slice2)
		Println(arr1[5])
	*/

	/*	var a, b int = 3, 4
		//f := sum

		f8(a, b, sum)
	*/
	/*	sum, sub := f9(10, 5)

		Println(sum, sub)

		f10(1, 2, 3, 4, 5, 6, 7, 8, 9)
	*/
}

/*指针传递*/
func f5(a *int) {
	*a = 666
}

/*数组名作为函数参数*/
func f6(a [5]int) {
	a[0] += 1
	Println(a[0])
}

/*Slice作为函数参数*/
func f7(a []int) {
	a[0] = 666
}

/*函数作为参数*/
func f8(a, b int, sum func(int, int) int) {
	Println(sum(a, b))
}

func sum(a, b int) int {
	return a + b
}

/*命名返回值参数,sum,sub*/
func f9(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

/*变参函数*/
func f10(args ...int) {
	for _, v := range args {
		Println(v)
	}
}

/*变参的传递*/
