package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//write
//Read

func create() {
	file1, _ := os.Create("hello1.txt")
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "Hello,world!")

	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 1, 1.1, "Hello,World!")

	file3, _ := os.Create("hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "hello,World!")
}
func read() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("hello1.txt")
	defer file1.Close()
	n, _ := fmt.Fscan(file1, &num1, &num2, &s)

	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2, s)

	file2, _ := os.Open("hello2.txt")
	defer file2.Close()
	fmt.Fscanln(file2, &num1, &num2, &s)
	fmt.Println(num1, num2, s)

	file3, _ := os.Open("hello3.txt")
	defer file3.Close()
	fmt.Fscan(file3, "%d,%f,%s", &num1, &num2, &s)
	fmt.Println(num1, num2, s)
}

func main() {
	create()
	read()
	////////// Create Close
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	s := "Hello, world"

	n, err := file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
	file.Close()
	///////// Read

	file, err = os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())
	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))

	file.Close()
	/////////
	file, err = os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDONLY|os.O_TRUNC, //읽기/쓰기, 파일이 있다면 연 뒤 내용 삭제

		os.FileMode(0644)) //파일권한은 644
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n = 0
	s = "안녕하세요"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장완료")

	fi, err = file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	data = make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)

	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))

	////// ioutil

	s = "hello, world"

	err = ioutil.WriteFile("hello.txt", []byte(s), os.FileMode(644))

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err = ioutil.ReadFile("hello.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}
