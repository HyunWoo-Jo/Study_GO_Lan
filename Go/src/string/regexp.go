package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("He", "Hello 100")
	fmt.Println(matched)
	matched, _ = regexp.MatchString("[a-zA-Z0-9+", "Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^Hello", "Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^Hello", "Go Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("Hello$", "Go Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("Hello$", "Go Hello")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Go.Hello(100)")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Go.Hello(100).x")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Go.Hello()")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\[a-zA-Z0-9]+\\[@]+\\[a-zA-Z].\\[a-zA-Z]$", "ASD212@asd.com")
	fmt.Println(matched)
	//////////////
	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)

	fmt.Println(validEmail.MatchString(("hello@mail.example.com")))
	/////////////
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.FindString("hello example.com")
	fmt.Println(s1)

	re2, _ := regexp.Compile("([a-zA-Z]+)(\\.[a-zA-Z]+)$")
	s2 := re2.FindStringSubmatch("hello example.com")

	fmt.Println(s2)

	n2 := re2.FindStringSubmatchIndex("example.com")

	fmt.Println(n2)

	re3, _ := regexp.Compile("\\.[a-zA-Z]+")
	s3 := re3.FindAllString("hello.co.kr world hello.net example.com", -1)

	fmt.Println(s3)

	n3 := re3.FindAllStringIndex("hello.co.kr wrold hello.net example.com", -1)

	fmt.Println(n3)
	/////////

	s1 = re1.ReplaceAllString("hello example.com", ".net")
	fmt.Println(s1)

}
