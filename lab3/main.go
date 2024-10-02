package main

import (
	"fmt"
	"ppLABS/lab3/mathutils"
	"ppLABS/lab3/stringutils"
)

func main() {
	fmt.Println(mathutils.Factorial(5))

	var someNum int
	fmt.Scan(&someNum)
	fmt.Println(mathutils.Factorial(someNum))

	stringutils.ReverseString("KokoJAMBO")

	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	array := [...]int{5, 4, 6, 7, 3}

	newSlice := array[:]

	newSlice[2] = 901
	fmt.Println(array, newSlice)

	newSlice = append(newSlice, 20)

	newSlice[1] = 89
	fmt.Println(array, newSlice)

	newSlice = newSlice[:len(newSlice)-1]
	fmt.Println(newSlice)

	strSlice("KokoJAMBO", "эх", "страно, что вывело меня да?", ":3")
}

func strSlice(strs ...string) {
	strArr := make([]string, len(strs))
	copy(strArr, strs)

	max := ""
	for _, v := range strArr {
		if utf8.RuneCountInString(v) > utf8.RuneCountInString(max) {
			fmt.Println(utf8.RuneCountInString(v))
			max = v
		}
	}

	fmt.Println(max)
}