package main

import (
	"os"

	//"os"
	"bufio"
	"bytes"
	. "fmt"

	//"io"
	"io/ioutil"
	//"os"
)

func main1() {
	/*
		slice1 := make([]string, 10)
		defer func() {
			for _, name := range slice1 {
				os.Remove(name)
			}
		}()

		for i := 0; i < 10; i++ {
			name, err := ioutil.TempDir("./", "gotmp")
			if err != nil {
				Println(err.Error())
				return
			}
			Println("dir:", name)
			slice1 = append(slice1, name)
		}
	*/

	/*
		slice1 := make([]*os.File, 10)

		for i := 0; i < 10; i++ {
			name, err := ioutil.TempFile("./", "")
			if err != nil {
				Println(err.Error())
				return
			}
			Println("file:", name.Name())
			slice1 = append(slice1, name)
		}

		func() {
			for _, name := range slice1 {
				name.Close()
				os.Remove(name.Name())
			}
		}()
	*/

}

func main2() {
	/*
		var (
			name string
			age  int
		)
	*/

	/*
		n1, _ := Sscan("polaris\n18", &name, &age)
		Println(n1, name, age)
		n2, _ := Sscan("polaris 18", &name, &age)
		Println(n2, name, age)
	*/

	/*
		n1, _ := Sscanf("polaris\n18", "%s%d", &name, &age)
		Println(n1, name, age)
		n2, _ := Sscanf("polaris 18", "%s%d", &name, &age)
		Println(n2, name, age)
	*/

	//	Scanf("%s%d", &name, &age)
	//Scanln(&name, &age)
	//Println(name, age)

	/*
		age++

		Scanf("%s\n", &name)
		Println(name)
		Scanf("%s\n", &name)
		Println(name)
	*/

	/*
		Scanln(&name)
		Println(name)
		Scanln(&age)
		Println(age)
	*/
}

func main() {
	/*
		reader1 := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is thehome of gophers\n"))

		line, _ := reader1.ReadSlice('\n')
		Printf("the line :%s\n", line)

		n, _ := reader1.ReadSlice('\n')
		Printf("the line :%s\n", line)
		Println(string(n))
	*/

	/*利用Scanner的split方法和Scan方法意打印出有多少单词*/
	/*	const input = "T h i s   is The Golang Standard Library.\nWelcome you!"
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanWords)

		count := 0

		for scanner.Scan() {
			Println(scanner.Text())
			count++
		}

		if err := scanner.Err(); err != nil {
			Fprintln(os.Stderr, "reading input:", err)
		}

		Println("count:", count)
	*/

	/*利用Scanner读取文件中数据，一次读取一行*/
	/*
		reader1 := bytes.NewBuffer(nil)
		reader1.ReadFrom(os.Stdin)

		ioutil.WriteFile("./1.txt", reader1.Bytes(), 0666)

		buf, _ := ioutil.ReadFile("./1.txt")
		scanner := bufio.NewScanner(bytes.NewReader(buf))

		scanner.Split(bufio.ScanLines)

		//输出

		for scanner.Scan() {
			Println(scanner.Text())
		}

		//
		if err := scanner.Err(); err != nil {
			Fprintln(os.Stderr, "Read error:", err)
		}
	*/
}
