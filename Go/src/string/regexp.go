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

	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)

	fmt.Println(validEmail.MatchString(("hello@mail.example.com")))
}
