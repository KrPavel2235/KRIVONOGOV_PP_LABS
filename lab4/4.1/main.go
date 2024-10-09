package main

import (
	"fmt"
)

func main() {
	//1 номер
	peopleMap := map[string]int{
		"Джексон Панин": 57,
		"Жэкич Злой":    21,
		"Артас Папаня":  45,
		"Хумас Ржачков": 24,
	}
	
	fmt.Println(peopleMap)

	addPerson(&peopleMap, "Павлик Кривой", 20)

	fmt.Println(peopleMap)

}

func addPerson(peopleMap *map[string]int, name string, age int) {
	(*peopleMap)[name] = age
}


