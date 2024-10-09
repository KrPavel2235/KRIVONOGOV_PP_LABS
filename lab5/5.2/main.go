package main

import (
	"fmt"
)

type Person struct {
	firstName  string
	lastName   string
	middleName string
	age        int
}

func (p Person) getPersonInfo() {
	fmt.Println("Имя: ", p.firstName)
	fmt.Println("Фамилия: ", p.lastName)
	fmt.Println("Отчество: ", p.middleName)
	fmt.Println("Возраст: ", p.age)
}

func (p *Person) birthday() {
	p.age++
}

func main() {
	firstPerson := Person{
		"Дубровин",
		"Руслан",
		"Владимирович",
		20,
	}
	firstPerson.getPersonInfo()

	fmt.Println("	Опа днюха:	")
	firstPerson.birthday()
	firstPerson.getPersonInfo()
}
