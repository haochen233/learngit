// chapter9 project chapter9.go
package main

import (
	. "fmt"
	"reflect"
)

type People struct {
	Name string
}

type Student struct {
	People
	class string
}

type Teacher struct {
	People
	WorkType string
}

//定义接口Speaker和Learner
type Speaker interface {
	Speak()
}

type Learner interface {
	Speaker
	Learn()
}

func (p People) Speak() {
	Println("I am a people")
}

func (s Student) Speak() {
	Println("I am a student")
}

/*
func (t Teacher) Speak() {
	Println("I am a teacher")
}
*/

func (s Student) Learn() {
	Println("I am a Student ,I like to study")
}

func main() {
	/*
		f1(1)
		f2("eq3reqwew")
		f2(true)
	*/

	//f2(1, 3, 4, '2', "daswdas", true, false)

	/*
		ret := f3()
		Println(ret)
	*/
	/*
		people := People{"张三"}
		teacher := Teacher{People{"李丽"}, "英语"}
		student := Student{People{"李四"}, "软件"}

		var is Speaker

		is = people
		is.Speak()

		is = teacher
		is.Speak()

		is = student
		is.Speak()

		var il Learner

		il = student
		il.Learn()

		ix := make([]Speaker, 3)
		ix[0], ix[1], ix[2] = people, student, teacher
		for i, _ := range ix {
			ix[i].Speak()
		}
	*/

	/*	peo := People{"张三"}

		tea := Teacher{People{"李莉"},
			"English",
		}

		var is Speaker

		is = &peo
		is.Speak()

		is = &tea
		is.Speak()
	*/

	/*接口的转换*/
	/*
		var id DerivedInfoer = Derived{Base{"张晨昊", 21},
			"信息16",
		}

		id.DerivedInfo()
		id.BaseInfo()

		var ib BaseInfoer
		ib = id
		ib.BaseInfo()
	*/

	/*接口类型推断——Comma-ok 断言*/
	/*	people := People{"张三"}
		teacher := Teacher{People{"李丽"}, "英语"}
		student := Student{People{"李四"}, "软件"}
		ix := make([]Speaker, 3)
		ix[0], ix[1], ix[2] = &people, student, &teacher
		for i, _ := range ix {
			ix[i].Speak()
		}

		for i, v := range ix {

			if _, ok := v.(*People); ok {
				Printf("ix[%d] : %v, It is a type *People\n", i, v)
			} else if _, ok := v.(Student); ok {
				Printf("ix[%d] : %v, It is a type Student\n", i, v)
			} else if _, ok := v.(*Teacher); ok {
				Printf("ix[%d] : %v, It is a type Teacher\n", i, v)
			}

		}
	*/

	/*接口类型推断——Switch 测试*/
	/*	people := People{"张三"}
		teacher := Teacher{People{"李丽"}, "英语"}
		student := Student{People{"李四"}, "软件"}
		ix := make([]Speaker, 3)
		ix[0], ix[1], ix[2] = &people, student, &teacher
		for i, _ := range ix {
			ix[i].Speak()
		}

		for i, v := range ix {
			switch value := v.(type) {
			case People:
				Printf("ix[%d] : %v, It is a type People %v\n", i, v, value)
			case *People:
				Printf("ix[%d] : %v, It is a type People %v\n", i, v, value)
			case Student:
				Printf("ix[%d] : %v, It is a type Student\n", i, v)
			case *Student:
				Printf("ix[%d] : %v, It is a type *Student\n", i, v)
			case Teacher:
				Printf("ix[%d] : %v, It is a type Teacher\n", i, v)
			case *Teacher:
				Printf("ix[%d] : %v, It is a type *Teacher\n", i, v)
			}
		}
	*/

	/*获取原对象的Type和Value值*/
	/*
		var pi float32 = 3.14

		t := reflect.TypeOf(pi)
		v := reflect.ValueOf(pi)

		var stu *Student

		t1 := reflect.TypeOf(stu)

		f := func() {
			//...
		}
		t2 := reflect.TypeOf(f)

		ff := f2
		t3 := reflect.TypeOf(ff)

		Println(t.Kind())
		Println(t, v)
		Println(t1)
		Println(t2)
		Println(t3)
	*/
	/*
		var person Person = Person{"张三",
			17,
			true,
		}

		Println(person)

		SetValue(&person)

		Println(person)
	*/
	/*
		per := Person{"张晨昊",
			21,
			false,
		}

		v := reflect.ValueOf(per)

		mv := v.MethodByName("Say")

		args := []reflect.Value{reflect.ValueOf("English"), 12, true}

		mv.Call(args)
	*/
	/*
		var p = &Person{"张晨昊",
			21,
			false,
		}

		typeofp := reflect.TypeOf(p)
		Printf("name :'%v', kind : '%v' \n", typeofp.Name(), typeofp.Kind())

		typeofp = typeofp.Elem()

		Printf("element name : '%v', Element kind: '%v'\n", typeofp.Name(),
			+typeofp.Kind())
	*/

	/*通过类型信息创建实例*/
	/*	var a int
		typeofa := reflect.TypeOf(a)
		ains := reflect.New(typeofa)

		Println(ains.Type(), ains.Kind())
	*/

}

type Person struct {
	Name string
	Age  int
	Sex  bool
}

func (p Person) Say(str string) {
	Printf("I am %s, I like %s\n", p.Name, str)
}

/*反射修改原对象Value值*/
/*
func SetValue(src interface{}) {
	v := reflect.ValueOf(src)

	if v.Kind() != reflect.Ptr {
		Println("cannot set")
		return
	} else {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString("张晨昊")
		case reflect.Int:
			v.Field(i).SetInt(21)
		case reflect.Bool:
			v.Field(i).SetBool(false)
		}
	}



}
*/

/*
type Base struct {
	Name string
	Age  int
}

type Derived struct {
	Base
	class string
}

type BaseInfoer interface {
	BaseInfo()
}

type DerivedInfoer interface {
	BaseInfo()
	DerivedInfo()
}

func (b Base) BaseInfo() {
	Println(b)
}

func (d Derived) DerivedInfo() {
	Println(d)
}



*/

/*空接口*/

//接收一个任意类型参数
func f1(a interface{}) {
	Println("this is f1,接收一个任意类型参数")
}

//接收任意多个类型参数
func f2(args ...interface{}) {
	Println("this is f2,接收多个任意类型参数")

}

//返回任意类型参数
func f3() interface{} {

	Println("this is f3,返回一个任意类型参数")
	return "中国"
}
