// chapter6 project main.go
package main

import (
	"bytes"
	. "fmt"
)

func main() {

	/*数组的定义*/

	/*
		var arr1 [5]int

		for i := 0; i < 5; i++ {
			arr1[i] = i + 1
		}

		Println(arr1)
	*/

	/*数组初始化形式种类*/

	/*	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		var b = [10]int{1, 2, 3, 4, 5}
		//[]必须加...否则就是创建的切片
		var c = [...]int{6, 7, 8, 9, 10, 11, 12, 13}


		var d = [10]int{1: 1, 3: 1, 5: 1, 7: 1, 9: 1}
		Println(a)
		Println(b)
		Println(c)
		Println(d)
	*/

	/*打印斐波那契数列的前20项，每行输出5项*/
	/*
		var fibo = [20]int{0: 1, 1: 1}

		for i := 2; i < 20; i++ {
			fibo[i] = fibo[i-2] + fibo[i-1]
		}

		for i := 0; i < 20; i++ {
			if i%5 == 0 {
				Println("")
			}

			Printf("fibo[%2d]=%-4d ", i, fibo[i])
		}
	*/
	/*用 for range 也可以遍历*/
	/*
		for i, v := range fibo {
			if i%5 == 0 {
				Println()
			}

			Printf("fibo[%2d]=%-4d ", i, v)
		}
	*/

	/*多维数组*/
	/*
		var a [3][4]int
		var b = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
		var c = [][]int{{3, 4, 5}, {7, 8}}
		var d = [2][2]int{{1: 1}, {1: 2}}

		Println(a)
		Println(b)
		Println(c)
		Println(d)
	*/

	/*基于底层数组创建切片*/

	/*	//先定义并初始化底层数组
		var array1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		//注意切片类型与底层数组类型保持一致
		var slice1 []int
		//部分引用的三种形式

			slice1 = array1[:5]
			slice2 := array1[5:]
			slice3 := array1[4:7]

			//全部引用的四种形式
			slice4 := array1[0:10]
			slice5 := array1[0:len(array1)]
			slice6 := array1
			slice7 := array1[:]

			Println(slice1)
			Println(slice2)
			Println(slice3)
			Println(slice4)
			Println(slice5)
			Println(slice6)
			Println(slice7)
	*/

	/*直接和使用make创建切片*/

	/*
		//直接创建
		var slice1 = []int{1, 2, 3, 4, 5}
		//注意：一定要与数组初始化 var array1 = [...]int{1,2,3}区分开来
		//使用make创建
		var slice2 = make([]int, 5)

		//使用make创建并且预留元素存储空间
		var slice3 = make([]int, 5, 10)

		//输出它们的元素个数size与容量capicity
		Printf("slice1_len = %d, capicity = %2d %v\n", len(slice1), cap(slice1), slice1)
		Printf("slice2_len = %d, capicity = %2d %v\n", len(slice2), cap(slice2), slice2)
		Printf("slice3_len = %d, capicity = %2d %v\n", len(slice3), cap(slice3), slice3)
	*/

	/*切片元素的遍历*/

	/*
		var slice1 = []int{1, 2, 3, 4, 5}

		//用下标访问
		for i := 0; i < len(slice1); i++ {
			Printf("slice1[%d] = %d", i, slice1[i])
		}

		Println()

		//用range 访问
		for i, v := range slice1 {
			Printf("slice1[%d] = %d", i, v)
		}

		Println()
	*/

	/*字典的初始化和创建*/

	/*
		//声明未初始化，map1是nil
		var map1 map[string]int
		Println(map1)
		//使用make初始化
		var map2 map[string]int
		map2 = make(map[string]int)
		map2["张晨昊"] = 666
		Println(map2)

		//使用{}初始化
		var map3 = map[string]int{}

		map3["唐伯虎"] = 888
		Println(map3)
	*/

	/*字典项查找*/

	/*
		var map1 map[string]int
		map1 = make(map[string]int)

		map1["957"] = 666
		map1["369"] = 235

		//Key存在的情况

		//delete(map1, "369")

		v, OK := map1["369"]
		if OK {
			Println(v, OK)
		} else {
			Println("不存在key值369")
		}

		//Key不存在的情况
		v1, OK := map1["233"]
		if OK {
			Println(v1, OK)
		} else {
			Println("不存在key值233")
		}

	*/

	/*new和make的使用对比*/

	//使用new为切片slice1分配内存

	/*	var p *[]int = new([]int)

		Printf("%d\n", p)
		Printf("%d\n", *p)

		//var slice1 = make([]int, 5, 10)
		var slice2 []int
		slice2 = make([]int, 5, 10)
		Println(slice2)
	*/

	/*Write() 写入缓冲函数test*/
	/*	p := []byte("Hello world!")

		var slice1 = make([]byte, 0, 10)

		Println(cap(slice1))

		b := bytes.NewBuffer(slice1)
		n, err := b.Write(p)

		Println(string(b.Bytes()), n, err, b.Cap())
	*/

	/*WriteByte() test*/
	/*	var c1, c2 byte = 'G', 'o'

		buffer1 := bytes.NewBuffer(nil)

		err := buffer1.WriteByte(c1)

		buffer1.WriteByte(c2)
		Println(buffer1, err)
	*/

	/*WriteRune() test*/
	/*	var r1, r2 rune = '中', '国'

		b := bytes.NewBuffer(nil)

		n, err := b.WriteRune(r1)
		b.WriteRune(r2)

		Println(string(b.Bytes()), n, err)
	*/

	/*WriteString() test*/
	/*	var str1 string = "Hello 中国"

		b := bytes.NewBuffer(nil)

		n, erro := b.WriteString(str1)

		Println(string(b.Bytes()), n, erro)
	*/

	/*WriteTo() test*/
	/*	buf := []byte("wo jiu shi Tian")

		b := bytes.NewBuffer(buf)
		w := bytes.NewBuffer(nil)

		n, err := b.WriteTo(w)

		Println(string(w.Bytes()), n, err)
	*/

	/*Read() test*/

}
