// 런타임에 인터페이스나 구조체 들의 타입 정보를 얻어내거나 결정하는 기능
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func main() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind() == reflect.Float64)

	fmt.Println(t.Kind() == reflect.Int64)

	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Kind() == reflect.Int64)
	fmt.Println(v.Float())

	p := Person{}
	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	//fmt.Println(reflect.ValueOf(a).int()) 런타임 에러
	fmt.Println(reflect.ValueOf(a).Elem())
	fmt.Println(reflect.ValueOf(a).Elem().Int())
}
