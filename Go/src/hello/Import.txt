package main

import (
	"calc"
	"fmt"

	"github.com/golang/example/stringutil"
	//cmd에 폴더 경로로 들어와 go get 형식으로 패키지를 받아올수 있다.
	//cmd cd D:\Go_Workspace\Study_GO_Lan\Go\src\hello
	//go get
)

func main() {
	fmt.Println(stringutil.Reverse("olleh"))
	fmt.Println(calc.Sum(3, 4))
}
