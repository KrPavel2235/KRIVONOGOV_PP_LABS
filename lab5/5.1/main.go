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


func main() {
	//Номер 1
	firstPerson := Person{
		"Руслаааан",
		20,
	}
	firstPerson.getPersonInfo()

}
