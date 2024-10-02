package main

import (
	"fmt"
	"lab3/mathutils"
	"lab3/stringutils"
	"unicode/utf8"
)

func main() {
	fmt.Println("Факториал 5: ", mathutils.Factorial(5))

	var someNum int
	fmt.Scan(&someNum)
	fmt.Println("Факториал введённого числа: ", mathutils.Factorial(someNum))

	fmt.Println("Слово KokoJAMBO наоборот:")
	stringutils.ReverseString("KokoJAMBO")

	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	array := [...]int{5, 4, 6, 7, 3}
	fmt.Println(array)

	newSlice := array[:]

	newSlice[2] = 901
	
	newSlice[1] = 89

	newSlice = append(newSlice, 20)
	fmt.Println(array, newSlice)

	newSlice = newSlice[:len(newSlice)-2]
	fmt.Println("====")
	fmt.Println(newSlice)

	strSlice("KokoJAMBO", "эх", "страно, что вывело меня да?", ":3")
}

func strSlice(strs ...string) {
	strArr := make([]string, len(strs))
	copy(strArr, strs)

	max := ""
	for _, v := range strArr {
		if utf8.RuneCountInString(v) > utf8.RuneCountInString(max) {
			max = v
		}
	}

	fmt.Println(max)
}