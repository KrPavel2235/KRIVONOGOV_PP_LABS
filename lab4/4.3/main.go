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

	deleteFromMapByName(peopleMap, "Артас Папаня")

}

func deleteFromMapByName(mapExample map[string]int, str string) {
	fmt.Println("До удаления: ", mapExample)
	delete(mapExample, str)
	fmt.Println("После удаленя: ", mapExample)
}

