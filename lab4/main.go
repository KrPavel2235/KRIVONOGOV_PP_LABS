package main

import (
	"fmt"
	"strings"
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

	fmt.Println(peopleMap, peopleMap["Джексон Панин"])
	avgAge(peopleMap)

	deleteFromMapByName(peopleMap, "Артас Папаня")
	fmt.Println(peopleMap)

	//4 номер
	fmt.Println("Введите какое-нибудь слово, желательно с нижнем регистром)))")
	var str string
	fmt.Scan(&str)
	makeUppercase(str)


	fmt.Println("Введите несколько чисел (если введёте 0 магии прекратится)")
	var nums []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		nums = append(nums, input)
	}

	megaSum(nums)

	
	fmt.Println("Введите несколько чисел (Если введёте 0 волшебник обидется)")
	var nums2 []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		nums2 = append(nums2, input)
	}
	reverseArrOfNums(nums2)
}

func addPerson(peopleMap *map[string]int, name string, age int) {
	(*peopleMap)[name] = age
	fmt.Printf("Добавлен: %s, Возраст: %d\n", name, age)
}

func reverseArrOfNums(nums []int) {
	for _, v := range nums {
		defer fmt.Println(v)
	}
}

func megaSum(nums []int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}

func makeUppercase(str string) {
	fmt.Println(strings.ToUpper(str))
}

func deleteFromMapByName(mapExample map[string]int, str string) {
	fmt.Println("До удаления: ", mapExample)
	delete(mapExample, str)
	fmt.Println("После удаленя: ", mapExample)
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