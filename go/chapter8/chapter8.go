// chapter8
package main

import (
	. "fmt"
)

type student struct {
	id    int
	name  string
	sex   bool
	age   int
	class string
}

func main() {

	//用不同的方式创建结构体student变量
	/*	var stu1 student
		stu2 := student{}

		//为结构体变量赋值
		stu1.name = "唐伯虎"
		stu2.name = "祝枝山"
		stu1.age = 22
		stu2.age = 23

		Println(stu1)
		Println(stu2)
	*/

	/*	var p1 *student
		p1 = new(student)
		Println(p1)
	*/

	/*结构体变量和对象的初始化*/
	/*
		//全部初始化
		stu1 := student{9527, "1", false, 20, "信息16"}
		stu2 := &student{9528, "2", true, 20, "信息16"}

		Println(stu1)
		Println(stu2)
		//部分初始化
		stu3 := student{name: "zhnag", age: 90}
		stu4 := &student{class: "软件16", name: "hou"}

		Println(stu3)
		Println(stu4)
	*/

	/*嵌入式结构直接定义结构体变量*/

	/*	var usr = struct {
			id   int
			name string
			age  int
		}{}

		usr2 := struct {
			id   int
			name string
			age  int
		}{9527, "张晨昊", 21}
		Println(usr)
		Println(usr2)
	*/

	/*	var map1 map[string]int

		map1 = make(map[string]int)

		map1["zhang"] = 61
		map1["liu"] = 62

		Println(map1)
	*/

	/*	map1 := map[string]struct {
			name string
			id   int
		}{
			"first":  {"zhnag", 20},
			"second": {"liu", 21},
		}

		Println(map1)
	*/
}
