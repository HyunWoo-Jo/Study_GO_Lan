package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	Write()
	// ioutil
	file, err := os.Open("hello.txt.gz")
	r1, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	r1.Close()

	b1, err := ioutil.ReadAll(r1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b1))
}
