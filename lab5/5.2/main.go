package main

import (
	"fmt"
)

type Person struct {
	Name  string
	age        int
}

func (p Person) getPersonInfo() {
	fmt.Println("Имя: ", p.Name)
	fmt.Println("Возраст: ", p.age)
}

func (p *Person) birthday() {
	p.age++
}

func main() {
	firstPerson := Person{
		"Гренка",
		20,
	}
	firstPerson.getPersonInfo()

	fmt.Println("	Опа днюха:	")
	firstPerson.birthday()
	firstPerson.getPersonInfo()
}
