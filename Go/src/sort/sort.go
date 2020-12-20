package main

import (
	"fmt"
	"sort"
)

//struct sort interface (len less swap)
type Student struct {
	name  string
	score float32
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByScore struct {
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score
}

//key default stucture
/*
type Data struct {
}

type By func(s1, s2 *Data) bool

func (by,By) Sort(data []Data) {
	sorter := &dataSorter{ // soter init
		data: data,
		by: by,
	}
	sort.Sort(sorter)
}

type dataSorter struct {
	data[]Data
	by func(s1, s2 *Data) bool
}

func (s *dataSorter) Len() int{
}
func (s *studentSorter) Less()(i,j int) bool {
}
func (s *studentSorter) Swap(i,j int){
}
*/
type By func(s1, s2 *Student) bool

func (by By) Sort(students []Student) {
	sorter := &studentSorter{
		students: students,
		by:       by,
	}
	sort.Sort(sorter)
}

type studentSorter struct {
	students []Student
	by       func(s1, s2 *Student) bool
}

func (s *studentSorter) Len() int {
	return len(s.students)
}
func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j])
}
func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func main() {
	a := []int{10, 5, 3, 7, 6}
	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	c := []string{"Maria", "Andrew", "John"}

	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	sort.Sort(sort.Float64Slice(b))
	fmt.Println(b)
	sort.Sort(sort.Reverse(sort.Float64Slice(b)))
	fmt.Println(b)

	sort.Sort(sort.StringSlice(c))
	fmt.Println(c)
	sort.Sort(sort.Reverse(sort.StringSlice(c)))
	fmt.Println(c)

	//struct

	st := Students{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}
	sort.Sort(st)
	fmt.Println(st)
	sort.Sort(sort.Reverse(ByScore{st}))
	fmt.Println(st)

	// key
	name := func(p1, p2 *Student) bool {
		return p1.name < p2.name
	}
	score := func(p1, p2 *Student) bool {
		return p1.score < p2.score
	}
	reverseScore := func(p1, p2 *Student) bool {
		return !score(p1, p2)
	}
	By(name).Sort(st)
	fmt.Println(st)

	By(reverseScore).Sort(st)
	fmt.Println(st)

}
