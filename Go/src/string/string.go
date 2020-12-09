package main

import (
	"fmt"
	"strconv"
)

func main1() {
	var s1 string
	s1 = strconv.FormatBool(true)
	fmt.Println(s1)

	s1 = strconv.FormatFloat(1.3, 'f', -1, 32)
	fmt.Println(s1)

	s1 = strconv.FormatInt(-10, 10)
	fmt.Println(s1)

	s1 = strconv.FormatUint(32, 16)
	fmt.Println(s1)

	var err error

	var num1 int
	num1, err = strconv.Atoi("100")
	fmt.Println(num1, err)

	var s []byte = make([]byte, 0)

	s = strconv.AppendBool(s, true)
	fmt.Println(string(s))

	s = strconv.AppendFloat(s, 1.3, 'f', -1, 32)
	fmt.Println(string(s))

	var b1 bool

	b1, err = strconv.ParseBool("true")
	fmt.Println(b1, err)
	var f1 float64
	f1, err = strconv.ParseFloat("1.3", 64)
	fmt.Println(f1, err)
}
