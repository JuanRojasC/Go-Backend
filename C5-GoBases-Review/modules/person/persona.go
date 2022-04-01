package person

import . "fmt"

func NewPerson(name string, age int) Person {
	return person{name, age}
}

type Person interface {
	GetAge() int
	Greeting()
}

type person struct {
	name string
	age  int
}

func (p person) GetAge() int {
	return p.age
}

func (p person) Greeting() {
	Printf("%s te saluda!!", p.name)
}
