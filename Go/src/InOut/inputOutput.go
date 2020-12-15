package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,

		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	w.WriteString("Hello, world!")
	w.Flush() // 버퍼 내용을 파일에 저장

	r := bufio.NewReader(file)

	fi, _ := file.Stat()

	b := make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)
	r.Read(b)

	fmt.Println(string(b))
	///////////////
	s1 := "Hello, world"
	r1 := strings.NewReader(s1)

	io.Copy(os.Stdout, r1)
	////////////
	r2 := bufio.NewReader(file)

	w2 := bufio.NewWriter(file)
	
	rw := bufio.NewReadWriter(r2, w2)

	rw.WriteString("Hello, world!")
	rw.Flush()
	
	file.Seek(0,os.SEEK_SET)

	data, _, _ := rw.ReadLine()
	fmt.Println(string(data))


}
