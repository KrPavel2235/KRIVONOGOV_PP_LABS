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

	addPerson(&peopleMap, "Павлик Кривой", 20)

	avgAge(peopleMap)

}

func addPerson(peopleMap *map[string]int, name string, age int) {
	(*peopleMap)[name] = age
}

func avgAge(mapExample map[string]int) {
	ages := 0
	counter := 0
	for _, v := range mapExample {
		ages += v
		counter++
	}
	fmt.Println(float64(ages) / float64(counter))
}
