package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("꿱")
}
func (d Duck) feathers() {
	fmt.Println("오리는 흼색과 회색 털을 가지고 있습니다.")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("사람은 오리 흉내냅니다. 꽥~!")
}
func (p Person) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case Duck:
		return "Duck"
	}
	return ""
}

func main() {
	var donald Duck
	var john Person
	var quac Quacker

	inTheForest(donald)
	inTheForest(john)
	quac = donald

	fmt.Println(formatString(donald))

	if v, ok := quac.(Duck); ok {
		fmt.Println(v, ok)
	}

}
