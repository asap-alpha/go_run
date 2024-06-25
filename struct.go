package main

import "fmt"

type person struct {
	name string
	age  int
	history
}

type history struct {
	gender     string
	occupation string
	salary     int
}

func newPerson(name string, age int) *person {

	sp := person{name: name, age: age, history: history{gender: "male", salary: 220}}
	sp.age = 122
	return &sp
}

func main() {

	owner := newPerson("Asap", 23)
	fmt.Println("no pointer =>> ", owner)

	sp := owner

	sp.occupation = "Programmer"
	sp.age = 20
	sp.salary = 120

	fmt.Println("pointer ==>> ", sp)

	cat := struct {
		name  string
		color string
		age   int
	}{
		name:  sp.name,
		color: "Dart Brown",
		age:   sp.age,
	}

	fmt.Println("#d cat => ", cat)

}
