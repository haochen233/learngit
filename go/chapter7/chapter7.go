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

	//f11(1, 2, 3, 4, 5, 6, 7, 8, 9)

	//f13(2, "Go", 8, "language", 'a', false, 'A', 3.14)

	/*匿名函数的声明与调用*/

	//声明匿名函数并且赋给变量f
	/*	f := func(a, b int) int {
			return a + b
		}

		//对函数类型变量f进行调用
		sum := f(10, 56)
		Println(sum)

		//声明并直接调用匿名函数
		sum = func(a, b int) int {
			return a + b
		}(66, 600)

		Println(sum)
	*/

	/*函数的闭包*/
	/*
		f := closures(624)
		Println(f(1))
		Println(f(2))
	*/

	/*函数递归求斐波那契数列*/
	/*	var v int = 0

		for i := 0; i < 40; i++ {
			v = fibo(i)

			Printf("%-10d", v)
			if (i+1)%5 == 0 {
				Println()
			}
		}
	*/

	/*defer语句test*/
	/*	defer Println("I am first")
		Println("I am second")
	*/

	/*defer语句,逆序调用（FILO）*/

	/*	for i := 0; i < 5; i++ {
			defer Println(i)
		}
	*/

	//f14()
	//Println(f15())

	/*panic-and-recover机制*/
	/*	defer func() {
			Println("函数defer开始运行")
			err := recover()
			if err != nil {
				Println("程序异常退出：", err)
			} else {
				Println("程序正常退出")
			}
		}()

		f16(100)
	*/

}

/*panic*/
func f16(a int) {
	Println("函数f16开始运行")
	if a >= 100 {
		panic("参数超出范围!")
	} else {
		Println("函数f16调用结束")
	}
}

/*被延迟的匿名函数会读取外部函数的返回值，并且对外部函数的返回值赋值*/
func f15() (i int) {
	defer func() {
		i++
	}()
	i = 1
	return
}

/*defer语句声明时的变量的值，就是最后执行defer语句的值*/
func f14() {
	i := 0
	defer Println(i)
	i = 66
}

/*求斐波那契*/

func fibo(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

/*函数闭包,返回值是一个函数类型*/
func closures(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

/*任意类型变参*/

func f13(args ...interface{}) {
	var num_Int = make([]int, 0, 6)
	var num_String = make([]string, 0, 6)
	var num_Byte = make([]int32, 0, 6)
	var num_other = make([]interface{}, 0, 6)

	for _, arg := range args {

		switch v := arg.(type) {
		case int:
			num_Int = append(num_Int, v)
			break
		case string:
			num_String = append(num_String, v)
			break
		case int32:
			num_Byte = append(num_Byte, v)
			break
		default:
			num_other = append(num_other, v)
		}

	}

	Println(num_Int)
	Println(num_String)
	Println(num_Byte)
	Println(num_other)

}

/*声明f11、f12为变参函数*/
func f11(args ...int) {
	f12(args...)

	f12(args[:2]...)
}

func f12(args ...int) {
	Println(args)
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
