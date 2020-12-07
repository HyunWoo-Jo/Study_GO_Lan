package main

import (
	"fmt"
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
}
