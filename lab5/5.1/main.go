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


func main() {
	//Номер 1
	firstPerson := Person{
		"Дубровин",
		"Руслан",
		"Владимирович",
		20,
	}
	firstPerson.getPersonInfo()

}
