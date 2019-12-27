// chapter6 project main.go
package main

import (
	. "fmt"
	//	"io"
	"log"
	//"os"
	"path/filepath"
	//"time"
)

func main() {
	/*
		file, err1 := os.Open("doc.go") 	//Read-only
		checkout(err1)

		buf := make([]byte, 200)

		_, err2 := file.Read(buf)
		checkout(err2)

		Println(string(buf))

		_, err3 := file.Write([]byte("Println(Hello os.file"))
		checkout(err3)
	*/

	/*
		file, err1 := os.Open("doc.go")
		checkout(err1)

		buf1 := make([]byte, 5)
		buf2 := make([]byte, 10)

		file.Read(buf1)
		Println(string(buf1))
		file.ReadAt(buf2, 7)
		Println(string(buf2))
		file.Read(buf1)
		Println(string(buf1))
	*/

	/*Write写*/
	/*
		file, err1 := os.OpenFile("doc.go", os.O_RDWR|os.O_CREATE, 0666)
		checkout(err1)

		file.Seek(0, 2)

		_, err2 := file.Write([]byte("Println(\"hello os.File\")"))
		checkout(err2)
	*/

	/*
		file, err1 := os.Open("doc.go") //Read-only
		checkout(err1)

		buf := make([]byte, 200)

		_, err2 := file.Read(buf)
		checkout(err2)

		Println(string(buf))

		filestat, err3 := os.Stat("doc.go")
		checkout(err3)

		Println(filestat.Name())
		Println(filestat.Mode())
		Println(filestat.Size())
		Println(filestat.ModTime())

		//改变文件时间戳

		//	os.Chtimes("doc.go", time.Now(), time.Now())
		//	Println(filestat.ModTime())
		os.Remove("doc1.go")

		//err4 := os.Rename("1.txt", "doc1.go")
		//checkout(err4)

		file.REead

	*/

	/*
		s := filepath.Dir("d:/1/2/3")
		Println(s)

		s2 := filepath.Base("/1/2/1.txt")
		Println(s2)

		ext := filepath.Ext("/1/2/3/21321.go.txt")
		Println(ext)

		Println(filepath.IsAbs("d:/"))

		Println(filepath.Abs("doc.go"))

	*/

}

func checkout(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
