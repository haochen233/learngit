// chapter8
package main

import (
	. "fmt"
	"math"
	"reflect"
)

/*
type student struct {
	id    int
	name  string
	sex   bool
	age   int
	class string
}
*/

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

	/*匿名字段的嵌入*/

	/*	type people struct {
			name string
			age  int
			tall float32
		}

		type teacher struct {

			//嵌入了一个匿名字段people
			people

			level int
			tall  float32
		}

		type set struct {
			people
			int
			byte
		}

		var t1 = teacher{people{"李丽", 26, 1.75},
			3,
			2.01,
		}

		var t2 = &teacher{people{"小红", 28, 1.80},
			5,
			1.68,
		}


		Println(t1)
		Println(t2)

		Println(t1.tall)
		Println(t1.people.tall)
	*/

	/*Method 与 Method同名*/

	/*	rct1 := rectangle{3, 4}
		rct2 := rectangle{8, 9}

		Println(rct1.area(), rct2.area())

		Println(math.Pi)

		c1 := circle{2}
		Println(c1.area())
	*/

	/*Method的继承*/

	/*	student1 := student{people{"张晨昊", 22},
			"信息16",
		}
		teacher1 := teacher{people{"MT", 29},
			"yanan",
		}

		student1.say()
		teacher1.say()
	*/

	/*字段标签*/

	/*	type User struct {
			Id   int    "账号"
			Name string "姓名"
			Sex  bool   "性别"
		}

		u := User{26, "张晨昊", false}

		//使用TypeOf函数获取对象的类型

		t := reflect.TypeOf(u)

		//使用ValueOf获取对象的值

		v := reflect.ValueOf(u)

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			Printf("%s (%s=%v)\n", f.Tag, f.Name, v.Field(i).Interface())
		}
	*/

}

type people struct {
	name string
	age  int
}

type teacher struct {

	//嵌入了一个匿名字段people
	people
	address string
}

type student struct {
	people
	class string
}

func (p people) say() {
	Printf("I am %s.I am %d years old\n", p.name, p.age)
}

func (stu student) say() {
	Printf("I am %s.I am %d years old, 我是%s班的。\n", stu.name, stu.age, stu.class)
}

type rectangle struct {
	width  int
	height int
}

type circle struct {
	radius float64
}

func (recv circle) area() float64 {
	return recv.radius * recv.radius * 2 * math.Pi
}

func (recv rectangle) area() int {
	return recv.width * recv.height
}
