package main

import (
	"fmt"
	"ppLABS/lab3/mathutils"
	"ppLABS/lab3/stringutils"
)

func main() {
	//Номера 1-2
	fmt.Println(mathutils.Factorial(5))

	var someNum int
	fmt.Scan(&someNum)
	fmt.Println(mathutils.Factorial(someNum))

	//Номер 3
	stringutils.ReverseString("ZXC")

	//Номер 4
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	//Номер 5
	array := [...]int{5, 4, 6, 7, 3}

	newSlice := array[:]

	newSlice[2] = 777
	fmt.Println(array, newSlice)

	newSlice = append(newSlice, 322)

	newSlice[1] = 1488
	fmt.Println(array, newSlice)

	newSlice = newSlice[:len(newSlice)-1]
	fmt.Println(newSlice)

	//6 номер
	strSlice("три", "четыре", "пять букв", "очень много букв")
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