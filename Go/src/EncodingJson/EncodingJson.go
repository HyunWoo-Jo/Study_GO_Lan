package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Name  string `json:"name"` //소문자로 하고싶으면
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64
	Author  Author
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author
	Content    string
	Recommends []string
	Comments   []Comment
}

func main() {
	//// Unmarshal json -> go
	doc := `{
		"name": "maria",
		"age": 10
	}
	`
	var data map[string]interface{}
	json.Unmarshal([]byte(doc), &data)

	fmt.Println(data["name"], data["age"])

	///// strcut
	doc1 := `
	[{
		"Id": 1,
		"Title": "Hello, world!",
		"Author" : {
			"Name": "Maria",
			"Email": "maria@example.com"
		},
		"Content": "Hello~",
		"Recommends":[
			"John",
			"Andrew"
		],
		"Comments": [{
			"id": 1,
			"Author": {
				"Name": "Andrew",
				"Email": "andrew@hello.com"
			},
			"Content": "Hello Maria"
		}]
	}]
	`

	var data1 []Article

	json.Unmarshal([]byte(doc1), &data1)

	fmt.Println(data1)
	//// go -> json marshal
	data2 := make([]Article, 1)

	data2[0].Id = 1
	data2[0].Title = "Hello, World"
	data2[0].Author.Name = "Maria"
	data2[0].Author.Email = "maria@example.com"
	data2[0].Content = "hello"
	data2[0].Recommends = []string{"John", "Andrew"}
	data2[0].Comments = make([]Comment, 1)
	data2[0].Comments[0].Id = 1
	data2[0].Comments[0].Content = "Hello Maria"
	data2[0].Comments[0].Author.Email = "andrew@hello.com"
	data2[0].Comments[0].Author.Name = "Andrew"

	doc2, _ := json.Marshal(data2)

	fmt.Println(string(doc2))

	//// json file
	err := ioutil.WriteFile("./articles.json", doc2, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadFile("./articles.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	var data3 []Article

	json.Unmarshal(b, &data3)

	fmt.Println(data3)

}
